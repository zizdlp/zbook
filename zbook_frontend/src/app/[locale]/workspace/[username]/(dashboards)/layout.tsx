import UserSideBar from "@/components/sidebars/UserSideBar";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";
import { auth } from "@/auth";
import { redirect } from "@/navigation";
import { Suspense } from "react";
import UserSideBarLoading from "@/components/loadings/UserSideBarLoading";
import MainContentWrapper from "@/components/wrappers/MainContentWrapper";
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
      <div className="px-4 mx-auto max-w-[93rem] lg:px-8">
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
        <MainContentWrapper>{children}</MainContentWrapper>
      </div>
    );
  } else {
    redirect(`/auth/login`); // Navigate to the new post page
  }
}
