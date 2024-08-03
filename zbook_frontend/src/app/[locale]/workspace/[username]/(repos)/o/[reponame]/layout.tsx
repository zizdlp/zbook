import SideBarLayout from "@/components/sidebars/SideBarLayout";
import { Suspense } from "react";
import SideBarLoading from "@/components/loadings/SideBarLoading";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";
import FrpcSideBar from "@/app/[locale]/frpc/FrpcSideBar";

export default function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { reponame: string; username: string };
}) {
  return (
    <div className="px-4 mx-auto max-w-[92rem] lg:px-8">
      {/* <FrpcSideBar /> */}
      <Suspense fallback={<SideBarLoading />}>
        <SideBarLayout username={params.username} reponame={params.reponame} />
      </Suspense>
      <SideBarToggleSmall />
      {children}
    </div>
  );
}
