"use client";

import { Link } from "@/navigation";
import { useSearchParams } from "next/navigation";
import { usePathname } from "@/navigation";
import { useTranslations } from "next-intl";

export default function Pagination({
  totalPages,
  pageKey,
}: {
  totalPages: number;
  pageKey: string;
}) {
  const t = useTranslations("Pagination");
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const currentPage = Number(searchParams?.get(pageKey)) || 1;

  const createPageURL = (pageNumber: number | string) => {
    const params = new URLSearchParams(searchParams?.toString());
    params.set(pageKey, pageNumber.toString());
    return `${pathname}?${params.toString()}`;
  };
  function RenderPageLink({
    pageNumber,
    condition,
    label,
  }: {
    pageNumber: number;
    condition: boolean;
    label?: string;
  }) {
    return (
      <Link
        key={pageNumber}
        href={createPageURL(pageNumber)}
        className={`mx-0.5 cursor-pointer ${isHidden(
          condition
        )} hover:text-sky-500`}
      >
        {label || pageNumber}
      </Link>
    );
  }
  const isHidden = (condition: boolean) => (condition ? "hidden" : "");

  return (
    <div className="flex items-center justify-start bg-transparent  py-1 ">
      <div className="items-start justify-start ">
        <span className="text-sm">
          <span className="mr-2">
            {t("TotalPage", { duration: totalPages })}
          </span>
          {totalPages > 1 && (
            <>
              <RenderPageLink
                pageNumber={currentPage - 1}
                condition={currentPage == 1}
                label={t("Previous")}
              />
              <RenderPageLink pageNumber={1} condition={currentPage == 1} />
              <span className={`${isHidden(currentPage - 2 <= 1)}`}>
                {t("Ellipsis")}
              </span>
              <RenderPageLink
                pageNumber={currentPage - 2}
                condition={currentPage - 2 <= 1}
              />
              <RenderPageLink
                pageNumber={currentPage - 1}
                condition={currentPage - 1 <= 1}
              />
              <span className="mx-0.5 cursor-pointer text-sky-500">
                {currentPage}
              </span>
              <RenderPageLink
                pageNumber={currentPage + 1}
                condition={currentPage + 1 >= totalPages}
              />
              <RenderPageLink
                pageNumber={currentPage + 2}
                condition={currentPage + 2 >= totalPages}
              />
              <span className={`${isHidden(currentPage + 3 >= totalPages)}`}>
                {t("Ellipsis")}
              </span>
              <RenderPageLink
                pageNumber={totalPages}
                condition={currentPage == totalPages}
              />
              <RenderPageLink
                pageNumber={currentPage + 1}
                condition={currentPage == totalPages}
                label={t("Next")}
              />
            </>
          )}
        </span>
      </div>
    </div>
  );
}
