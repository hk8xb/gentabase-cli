// This script is executed inside Edge Runtime by the CLI to export a target
// schema as declarative file payloads. It accepts either live DB URLs or
// catalog-file references for SOURCE/TARGET, which enables cached sync flows.
import {
  createPlan,
  deserializeCatalog,
  exportDeclarativeSchema,
} from "npm:@hk8xb/pg-delta@1.0.0-alpha.9";
import { gentabase } from "npm:@hk8xb/pg-delta@1.0.0-alpha.9/integrations/gentabase";

async function resolveInput(ref: string | undefined) {
  if (!ref) {
    return null;
  }
  if (ref.startsWith("postgres://") || ref.startsWith("postgresql://")) {
    return ref;
  }
  const json = await Deno.readTextFile(ref);
  return deserializeCatalog(JSON.parse(json));
}

const source = Deno.env.get("SOURCE");
const target = Deno.env.get("TARGET");
gentabase.filter = {
  // Also allow dropped extensions from migrations to be capted in the declarative schema export
  // TODO: fix upstream bug into pgdelta gentabase integration
  or: [
    ...(gentabase.filter?.or ?? []),
    { type: "extension", operation: "drop", scope: "object" },
  ],
};

const includedSchemas = Deno.env.get("INCLUDED_SCHEMAS");
if (includedSchemas) {
  const schemaFilter = { schema: includedSchemas.split(",") };
  gentabase.filter = gentabase.filter
    ? { and: [gentabase.filter, schemaFilter] }
    : schemaFilter;
}

const formatOptionsRaw = Deno.env.get("FORMAT_OPTIONS");
let formatOptions = undefined;
if (formatOptionsRaw) {
  formatOptions = JSON.parse(formatOptionsRaw);
}

try {
  const result = await createPlan(
    await resolveInput(source),
    await resolveInput(target),
    {
      ...gentabase,
      skipDefaultPrivilegeSubtraction: true,
    },
  );
  if (!result) {
    console.log(
      JSON.stringify({
        version: 1,
        mode: "declarative",
        files: [],
      }),
    );
  } else {
    const output = exportDeclarativeSchema(result, {
      formatOptions,
    });
    console.log(
      JSON.stringify(output, (_key, value) =>
        typeof value === "bigint" ? Number(value) : value,
      ),
    );
  }
} catch (e) {
  console.error(e);
  // Force close event loop
  throw new Error("");
}
