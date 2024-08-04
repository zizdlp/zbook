import MainContentWrapper from "../wrappers/MainContentWrapper";
import ContentBarLoading from "./ContentBarLoading";
import LoadingElement from "./LoadingElement";

export default function MarkdownLoading() {
  return (
    <MainContentWrapper right_sidebar={<ContentBarLoading />}>
      <div className="relative animate-pulse flex flex-col w-full">
        <LoadingElement className="rounded-md h-4 w-24 mb-2 " />
        <LoadingElement className="rounded-md h-6 w-1/3 mb-4 " />

        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-4" />
        <LoadingElement className="rounded-md h-96 mb-4" />

        <LoadingElement className="rounded-md h-5 w-48 mb-4" />
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-4" />
        <LoadingElement className="rounded-md h-96 mb-4" />

        <div className="flex items-center my-4">
          <LoadingElement className="rounded-full flex-none w-10 h-10 me-3" />
          <div>
            <LoadingElement className="rounded-md h-2.5 w-32 mb-2" />
            <LoadingElement className="rounded-md h-2 w-48" />
          </div>
        </div>
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-2.5" />
        <LoadingElement className="rounded-md h-4 mb-4" />
      </div>
    </MainContentWrapper>
  );
}
