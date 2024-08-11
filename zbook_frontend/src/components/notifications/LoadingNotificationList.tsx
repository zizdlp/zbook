import LoadingElement from "../loadings/LoadingElement";

export default function LoadingNotificationList({
  itemCount,
}: {
  itemCount: number;
}) {
  // 生成指定数量的占位符数组
  const placeholderItems = Array.from(
    { length: itemCount },
    (_, index) => index
  );

  return (
    <div>
      {placeholderItems.map((_, index) => (
        <div
          key={index}
          className="rounded-md md:rounded-xl dark:shadow-lg my-2 md:my-3  animate-pulse
          bg-white dark:bg-slate-800 hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] dark:border-0 flex items-center justify-between"
        >
          <div className="flex items-center justify-center">
            <div
              className={`flex-shrink-0 ml-4 w-2 h-2 md:w-2 md:h-2 rounded-full`}
            />
            <div className="flex-shrink-0">
              <LoadingElement className="w-8 h-8 md:w-12 md:h-12 ml-2 mr-4 rounded-full shadow-lg" />
            </div>
            <div className="flex flex-col md:py-5 py-3 space-y-2">
              <LoadingElement className="w-16 md:w-24 h-4 rounded-md" />
              <LoadingElement className="w-36 md:w-64 h-4 rounded-md" />
            </div>
          </div>
          <div className="px-4 flex-shrink-0">
            <LoadingElement className="w-16 h-6 rounded-md" />
          </div>
        </div>
      ))}
    </div>
  );
}
