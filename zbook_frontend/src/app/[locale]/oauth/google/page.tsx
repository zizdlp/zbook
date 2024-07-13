import LinkOAuth from "@/components/LinkOAuth";

export default async function GitTest({
  params,
  searchParams,
}: {
  params: { username: string };
  searchParams?: { access_token?: string };
}) {
  return <LinkOAuth oauthType="google" searchParams={searchParams} />;
}
