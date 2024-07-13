import UserSideBar from "@/components/sidebars/UserSideBar";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";
import ContentWrapper from "@/components/wrappers/ContentWrapper";
import { auth } from "@/auth";
import { redirect } from "@/navigation";
import { Suspense } from "react";
import UserSideBarLoading from "@/components/loadings/UserSideBarLoading";
export default async function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { username: string };
}) {
  const session = await auth();
  if (session && session.access_token) {
    return (
      <div className="overflow-hidden">
        <Suspense
          fallback={
            <UserSideBarLoading
              username={params.username}
              authname={session.username}
              authrole={session.role}
            />
          }
        >
          <UserSideBar
            username={params.username}
            authname={session.username}
            authrole={session.role}
          />
        </Suspense>
        <SideBarToggleSmall />
        <ContentWrapper>{children}</ContentWrapper>
      </div>
    );
  } else {
    redirect(`/auth/login`); // Navigate to the new post page
  }
}
