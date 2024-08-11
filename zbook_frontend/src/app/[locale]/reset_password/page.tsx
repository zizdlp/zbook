import ResetPasswordForm from "../auth/forget/ResetPasswordForm";

export default async function ResetPassword({
  searchParams,
}: {
  searchParams?: {
    verification_url?: string;
  };
}) {
  const verification_url = searchParams?.verification_url || "";
  return <ResetPasswordForm verification_url={verification_url} />;
}
