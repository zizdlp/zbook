import { Link } from "@/navigation";

import { CgHashtag } from "react-icons/cg";
import { BiChevronRight } from "react-icons/bi";
import parse, {
  domToReact,
  HTMLReactParserOptions,
  Element,
} from "html-react-parser";
import { MarkdownBasicInfo } from "@/fetchs/model";
interface ProfileProps {
  MarkdownBasicInfo: MarkdownBasicInfo;
}
export default function SearchMarkdownComponent(props: ProfileProps) {
  const options: HTMLReactParserOptions = {
    replace: (domNode) => {
      const typedDomNode = domNode as Element;
      if (typedDomNode.name === "b") {
        return (
          <b className="underline underline-offset-2">
            {domToReact(typedDomNode.children, options)}
          </b>
        );
      }
    },
  };
  const prefix = "/workspace/";
  return (
    <div
      className="rounded-md md:rounded-xl my-2 md:my-3 py-2 bg-white dark:bg-slate-700
      hover:dark:bg-sky-900 hover:bg-sky-500 hover:text-white border-[0.05rem] border-slate-300/75 dark:border-0 flex items-center justify-between"
    >
      <Link
        href={
          `${prefix}` +
          props.MarkdownBasicInfo.username +
          "/o/" +
          props.MarkdownBasicInfo.repo_name +
          "/" +
          props.MarkdownBasicInfo.relative_path
        }
        className="flex items-center justify-between w-full px-2"
      >
        <div className="flex items-center justify-center">
          <div className="p-2">
            <CgHashtag className="w-5 h-5" />
          </div>
          <div className="flex-grow overflow-auto">
            <span>{props.MarkdownBasicInfo.relative_path}</span>
            <br />
            <span>
              {parse(props.MarkdownBasicInfo?.main_content ?? "", options)}
            </span>
          </div>
        </div>

        <div className="p-2">
          <BiChevronRight className="w-5 h-5" />
        </div>
      </Link>
    </div>
  );
}
