import LoadingElement from "@/components/loadings/LoadingElement";
import SideBarLiContent from "@/components/sidebars/SideBarLiContent";
export default function GroupLoading({
  itemCount,
  showRight,
}: {
  itemCount: number;
  showRight: boolean;
}) {
  const placeholderItems = Array.from(
    { length: itemCount },
    (_, index) => index
  );
  return (
    <div className="w-full pt-4 rounded-md sm:pt-6">
      <div className="font-semibold text-slate-800 dark:text-slate-400 border-slate-300 dark:border-slate-700 border-b-[0.01rem] pb-0.5">
        <LoadingElement className="w-16 h-5 rounded-md animate-pulse mb-0.5" />
      </div>
      <ul className="my-4 space-y-3 animate-pulse">
        {placeholderItems.map((_, index) => (
          <li key={index}>
            <SideBarLiContent isSelected={false}>
              <div className="flex items-center justify-between w-full">
                <div className="flex items-center justify-center">
                  <LoadingElement className="h-6 w-6 rounded-md" />
                  <LoadingElement className="h-6 w-12  ms-3 rounded-md" />
                </div>
                {showRight && (
                  <LoadingElement className="inline-flex items-center ms-3 justify-center h-6 w-6 rounded-md" />
                )}
              </div>
            </SideBarLiContent>
          </li>
        ))}
      </ul>
    </div>
  );
}
