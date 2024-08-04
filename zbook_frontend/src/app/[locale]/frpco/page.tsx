import FrpcSideBar from "../frpc/FrpcSideBar";

export default function FrpcO() {
  return (
    <div className="relative antialiased text-gray-500 dark:text-gray-400">
      <div className="px-4 mx-auto max-w-[92rem] lg:px-8">
        <FrpcSideBar />
        <div id="content-container">
          <div className="flex flex-row items-stretch gap-12 pt-[9.5rem] lg:pt-[6.5rem]">
            <div className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem] border">
              main
            </div>
            <div
              className="z-10 hidden xl:flex flex-none pl-10 w-[19rem] border"
              id="table-of-contents"
            >
              side
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
