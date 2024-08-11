import { Link } from "@/navigation";

import { CgHashtag } from "react-icons/cg";
import { BiChevronRight } from "react-icons/bi";
import parse, {
  domToReact,
  HTMLReactParserOptions,
  Element,
} from "html-react-parser";
import { MarkdownBasicInfo } from "@/fetchs/model";
import SearchItemWrapper from "./SearchItemWrapper";
import { useContext } from "react";
import { SideBarContext } from "@/providers/SideBarProvider";
interface ProfileProps {
  MarkdownBasicInfo: MarkdownBasicInfo;
}
export default function SearchMarkdownComponent(props: ProfileProps) {
  const { sideBarReload, setSideBarReload } = useContext(SideBarContext);
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
    <SearchItemWrapper>
      <div
        onClick={() => {
          console.log("refresh page");
          setSideBarReload(!sideBarReload);
          // refreshPage("/", true, false);
        }}
        className="w-full"
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
          <div className="flex items-center justify-center overflow-scroll">
            <div className="p-2">
              <CgHashtag className="w-5 h-5" />
            </div>
            <div className="flex-grow overflow-scroll">
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
    </SearchItemWrapper>
  );
}
