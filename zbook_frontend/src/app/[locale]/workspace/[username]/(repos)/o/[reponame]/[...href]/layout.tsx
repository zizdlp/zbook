import { Suspense } from "react";
import MarkdownLoading from "@/components/loadings/MarkdownLoading";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <> {children}</>;
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
