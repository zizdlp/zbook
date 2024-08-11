export default function LoadingList({ itemCount }: { itemCount: number }) {
  // 生成指定数量的占位符数组
  const placeholderItems = Array.from(
    { length: itemCount },
    (_, index) => index
  );

  return (
    <div className="grid lg:grid-cols-2 gap-4 grid-cols-1">
      {placeholderItems.map((_, index) => (
        <div
          key={index}
          className="border-[0.05rem] border-slate-300 dark:border-0 h-44 rounded-md p-4 flex flex-col
             dark:bg-slate-800/50 dark:hover:bg-slate-800 hover:bg-slate-50"
        >
          <div className="flex-none flex justify-between items-center">
            <div className="flex items-center justify-center space-x-2">
              <div className="flex-none w-12 h-12 rounded-full bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>
              <div>
                <div className="h-2.5 rounded-full bg-gray-200  dark:bg-gray-700/75 w-32 mb-2 animate-pulse"></div>
                <div className="w-48 h-2 rounded-full bg-gray-200  dark:bg-gray-700/75 animate-pulse"></div>
              </div>
            </div>
            <div className="flex-none w-12 h-6 rounded-md bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>
          </div>
          <div className="flex-1  my-1 rounded-md bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>

          <div className="flex-none flex justify-between items-center">
            <div className="flex-none w-12 h-6 rounded-md bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>
            <div className="flex items-center justify-center space-x-1">
              <div className="flex-none w-12 h-6 rounded-md bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>
              <div className="flex-none w-12 h-6 rounded-md bg-gray-200 dark:bg-gray-700/75 animate-pulse"></div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
