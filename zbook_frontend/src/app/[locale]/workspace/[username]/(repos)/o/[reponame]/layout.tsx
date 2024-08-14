import RepoSideBarLayout from "@/components/sidebars/RepoSideBarLayout";
import { Suspense } from "react";
import RepoSideBarLoading from "@/components/loadings/RepoSideBarLoading";
import SideBarToggleSmall from "@/components/sidebars/SideBarToggleSmall";
import MarkdownLoading from "@/components/loadings/MarkdownLoading";

export default function RootLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { reponame: string; username: string };
}) {
  return (
    <div className="px-4 mx-auto max-w-[93rem] lg:px-8">
      <Suspense fallback={<RepoSideBarLoading />}>
        <RepoSideBarLayout
          username={params.username}
          reponame={params.reponame}
        />
      </Suspense>
      <SideBarToggleSmall />
      <Suspense fallback={<MarkdownLoading />}>{children}</Suspense>
    </div>
  );
}
