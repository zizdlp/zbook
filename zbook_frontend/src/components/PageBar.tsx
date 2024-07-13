import { useTranslations } from "next-intl";

export default function PageBar({
  currentPage,
  totalPages,
  setCurrentPage,
}: {
  currentPage: number;
  totalPages: number;
  setCurrentPage: React.Dispatch<React.SetStateAction<number>>;
}) {
  const t = useTranslations("Pagination");
  const isHidden = (condition: boolean) => (condition ? "hidden" : "");

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
      <span
        onClick={() => {
          setCurrentPage(pageNumber);
        }}
        className={`mx-0.5 cursor-pointer ${isHidden(
          condition
        )} hover:text-sky-500`}
      >
        {label || pageNumber}
      </span>
    );
  }

  if (totalPages <= 1) {
    return <></>;
  }
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
          {/* <span
            onClick={() => {
              props.setCurrentPage(props.currentPage - 1);
            }}
            className={`${
              props.currentPage == 1 && "hidden"
            } cursor-pointer mr-0.5`}
          >
            {t("Previous")}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(1);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage == 1 && "hidden"
            }`}
          >
            1
          </span>
          <span className={`${props.currentPage - 2 <= 1 && "hidden"}`}>
            {t("Ellipsis")}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.currentPage - 2);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage - 2 <= 1 && "hidden"
            }`}
          >
            {props.currentPage - 2}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.currentPage - 1);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage - 1 <= 1 && "hidden"
            }`}
          >
            {props.currentPage - 1}
          </span>
          <span className="mx-0.5 cursor-pointer text-sky-500">
            {props.currentPage}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.currentPage + 1);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage + 1 >= props.pageNumber && "hidden"
            }`}
          >
            {props.currentPage + 1}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.currentPage + 1);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage + 2 >= props.pageNumber && "hidden"
            }`}
          >
            {props.currentPage + 2}
          </span>
          <span
            className={`${
              props.currentPage + 3 >= props.pageNumber && "hidden"
            }`}
          >
            {t("Ellipsis")}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.pageNumber);
            }}
            className={`mx-0.5 cursor-pointer ${
              props.currentPage == props.pageNumber && "hidden"
            }`}
          >
            {props.pageNumber}
          </span>
          <span
            onClick={() => {
              props.setCurrentPage(props.currentPage + 1);
            }}
            className={`${
              props.currentPage == props.pageNumber && "hidden"
            } cursor-pointer ml-0.5`}
          >
            {t("Next")}
          </span>{" "} */}
        </span>
      </div>
    </div>
  );
}
