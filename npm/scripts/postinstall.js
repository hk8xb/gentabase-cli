#!/usr/bin/env node
const { execSync } = require('child_process')
const fs = require('fs')
const path = require('path')
const https = require('https')
const os = require('os')
const crypto = require('crypto')
const { createGunzip } = require('zlib')
const tar = require('tar')

const VERSION = require('../package.json').version
const REPO = 'hk8xb/gentabase-cli'

function getPlatform() {
  const platform = os.platform()
  const arch = os.arch()

  const platformMap = {
    darwin: 'darwin',
    linux: 'linux',
    win32: 'windows',
  }

  const archMap = {
    x64: 'amd64',
    arm64: 'arm64',
  }

  const p = platformMap[platform]
  const a = archMap[arch]

  if (!p || !a) {
    throw new Error(`Unsupported platform: ${platform}/${arch}`)
  }

  return { platform: p, arch: a, ext: platform === 'win32' ? '.exe' : '' }
}

async function download(url) {
  return new Promise((resolve, reject) => {
    https.get(url, { headers: { 'User-Agent': 'gentabase-cli-npm' } }, (res) => {
      if (res.statusCode === 302 || res.statusCode === 301) {
        return download(res.headers.location).then(resolve, reject)
      }
      if (res.statusCode !== 200) {
        return reject(new Error(`Download failed: ${res.statusCode} ${url}`))
      }
      const chunks = []
      res.on('data', (chunk) => chunks.push(chunk))
      res.on('end', () => resolve(Buffer.concat(chunks)))
      res.on('error', reject)
    }).on('error', reject)
  })
}

async function main() {
  const { platform, arch, ext } = getPlatform()
  const binDir = path.join(__dirname, '..', 'bin')
  const binPath = path.join(binDir, `gentabase${ext}`)

  // Skip if binary already exists (e.g., from a previous install)
  if (fs.existsSync(binPath)) {
    console.log('Gentabase CLI binary already exists, skipping download.')
    return
  }

  const assetName = `gentabase_${platform}_${arch}.tar.gz`
  const url = `https://github.com/${REPO}/releases/download/v${VERSION}/${assetName}`

  console.log(`Installing Gentabase CLI v${VERSION} (${platform}/${arch})...`)

  try {
    const tarball = await download(url)

    // Extract the binary from the tarball
    fs.mkdirSync(binDir, { recursive: true })
    const tmpFile = path.join(os.tmpdir(), assetName)
    fs.writeFileSync(tmpFile, tarball)

    execSync(`tar xzf "${tmpFile}" -C "${binDir}" gentabase${ext}`, { stdio: 'pipe' })
    fs.chmodSync(binPath, 0o755)
    fs.unlinkSync(tmpFile)

    console.log('Gentabase CLI installed successfully.')
  } catch (err) {
    console.error(`Failed to install Gentabase CLI: ${err.message}`)
    console.error(`You can download it manually from: https://github.com/${REPO}/releases`)
    process.exit(0) // Don't fail npm install
  }
}

main()
