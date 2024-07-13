import { redirect } from "@/navigation";
import { auth } from "@/auth";
export default async function WorkSpacePage() {
  const session = await auth();
  if (session && session.access_token) {
    redirect(`/workspace/${session.username}`); // Navigate to the new post page
  } else {
    redirect(`/auth/login`); // Navigate to the new post page
  }
}
