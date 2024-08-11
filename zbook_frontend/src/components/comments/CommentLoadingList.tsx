import LoadingElement from "../loadings/LoadingElement";

export default function CommentLoadingList({
  itemCount,
}: {
  itemCount: number;
}) {
  const placeholderItems = Array.from(
    { length: itemCount },
    (_, index) => index
  );

  return (
    <div className="grid lg:grid-cols-1 gap-4 grid-cols-1">
      {placeholderItems.map((_, index) => (
        <div
          key={index}
          className="flex flex-col border-b dark:border-b-slate-700/75 border-b-slate-200/75 pt-4 text-slate-700 dark:text-slate-400 pb-2"
        >
          <div className="flex flex-row space-x-2 md:space-x-4">
            <LoadingElement className="h-10 w-10 rounded-full" />
            <div className="grow flex flex-col">
              <LoadingElement className="h-6 w-24 rounded-md my-0.5" />
              <div className="pt-0.5">
                <LoadingElement className="h-6 w-2/3 rounded-md my-0.5" />
                <div className="leading-6 mt-0.5 flex items-center justify-between">
                  <LoadingElement className="h-6 w-48 rounded-md my-0.5" />
                  <LoadingElement className="h-6 w-8 rounded-md my-0.5" />
                </div>
              </div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
