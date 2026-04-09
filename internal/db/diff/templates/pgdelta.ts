import {
  createPlan,
  deserializeCatalog,
  formatSqlStatements,
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

const includedSchemas = Deno.env.get("INCLUDED_SCHEMAS");
if (includedSchemas) {
  gentabase.filter = { schema: includedSchemas.split(",") };
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
    gentabase,
  );
  let statements = result?.plan.statements ?? [];
  if (formatOptions != null) {
    statements = formatSqlStatements(statements, formatOptions);
  }
  for (const sql of statements) {
    console.log(`${sql};`);
  }
} catch (e) {
  console.error(e);
  // Force close event loop
  throw new Error("");
}
