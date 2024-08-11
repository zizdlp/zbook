import { Suspense } from "react";
import MarkdownLoading from "@/components/loadings/MarkdownLoading";
import ContentBarLoading from "@/components/loadings/ContentBarLoading";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <Suspense
      fallback={
        <>
          <MarkdownLoading />
        </>
      }
    >
      {children}
    </Suspense>
  );
}
