import SideBarLayout from "@/components/sidebars/SideBarLayout";
import { Suspense } from "react";
import SideBarLoading from "@/components/loadings/SideBarLoading";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";

export default function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { reponame: string; username: string };
}) {
  return (
    <div>
      <Suspense fallback={<SideBarLoading />}>
        <SideBarLayout username={params.username} reponame={params.reponame} />
      </Suspense>
      <SideBarToggleSmall />
      {children}
    </div>
  );
}
