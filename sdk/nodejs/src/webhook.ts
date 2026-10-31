/** Webhook signature verification for HiTechCloud Sale AI. */
import { createHmac, timingSafeEqual } from "crypto";

export const SIGNATURE_HEADER = "X-SaleAI-Signature";

export function computeSignature(rawBody: string | Buffer, secret: string): string {
  const hmac = createHmac("sha256", secret);
  hmac.update(rawBody);
  return "sha256=" + hmac.digest("hex");
}

export function verifySignature(
  rawBody: string | Buffer,
  signature: string | undefined | null,
  secret: string,
): boolean {
  if (!signature) return false;
  const expected = computeSignature(rawBody, secret);
  const a = Buffer.from(expected);
  const b = Buffer.from(signature);
  if (a.length !== b.length) return false;
  return timingSafeEqual(a, b);
}
