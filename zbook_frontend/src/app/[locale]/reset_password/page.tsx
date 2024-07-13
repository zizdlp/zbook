import ResetPasswordForm from "../auth/forget/ResetPasswordForm";

export default async function ResetPassword({
  searchParams,
}: {
  searchParams?: {
    verification_id?: string;
  };
}) {
  const verification_id = searchParams?.verification_id || "";
  return <ResetPasswordForm verification_id={verification_id} />;
}
