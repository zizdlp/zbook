import SideBarLayout from "@/components/sidebars/SideBarLayout";
import { Suspense } from "react";
import LeftSideBarLoading from "@/components/loadings/LeftSideBarLoading";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";

export default function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { reponame: string; username: string };
}) {
  return (
    <div className="px-4 mx-auto max-w-[92rem] lg:px-8">
      <Suspense fallback={<LeftSideBarLoading />}>
        <SideBarLayout username={params.username} reponame={params.reponame} />
      </Suspense>
      <SideBarToggleSmall />
      {children}
    </div>
  );
}
