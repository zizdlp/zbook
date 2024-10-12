import { Element, domToReact } from "html-react-parser";
import { getClientAdmonitionType } from "@/utils/util";
import { motion } from "framer-motion";
import type { HTMLReactParserOptions } from "html-react-parser";
import { LuChevronLeftSquare, LuChevronDownSquare } from "react-icons/lu";
import { useState } from "react";
export default function Admonition({
  domNode,
  options,
}: {
  domNode: Element;
  options: HTMLReactParserOptions;
}) {
  let admtype = "note";
  let trigger = "!";
  const [isBodyVisible, setIsBodyVisible] = useState(true);

  const toggleBodyVisibility = () => {
    setIsBodyVisible((prev) => !prev);
  };

  const classAttr = domNode.attributes.find((attr) => attr.name === "class");
  const classList = classAttr ? classAttr.value.split(" ") : [];

  if (classList.includes("adm-note")) {
    admtype = "note";
  } else if (classList.includes("adm-warning")) {
    admtype = "warning";
  } else if (classList.includes("adm-info")) {
    admtype = "info";
  } else if (classList.includes("adm-tip")) {
    admtype = "tip";
  } else if (classList.includes("adm-error")) {
    admtype = "error";
  }

  const { bg, Icon } = getClientAdmonitionType(admtype);
  const admTitle = Array.from(domNode.childNodes).find((child) => {
    if (child instanceof Element && child.tagName.toLowerCase() === "div") {
      const classAttr = child.attributes.find((attr) => attr.name === "class");
      const classList = classAttr ? classAttr.value.split(" ") : [];
      if (classList.includes("admonition-?")) {
        trigger = "?";
      }
      return classList.includes("adm-title"); // Check if "adm-title" is in the class list
    }
    return false;
  });
  const admBody = Array.from(domNode.childNodes).find((child) => {
    if (child instanceof Element && child.tagName.toLowerCase() === "div") {
      const classAttr = child.attributes.find((attr) => attr.name === "class");
      const classList = classAttr ? classAttr.value.split(" ") : [];
      return classList.includes("adm-body"); // Check if "adm-title" is in the class list
    }
    return false;
  });
  if (
    admTitle &&
    admTitle instanceof Element &&
    admBody &&
    admBody instanceof Element
  ) {
    return (
      <div className="my-[1.25em] bg-slate-100/50 dark:bg-slate-800/50 ring-1 ring-slate-200/50 dark:ring-slate-900/10 rounded-md">
        <div
          className={`relative py-1 md:py-2 space-x-4 rounded-t-md flex items-center justify-center text-slate-400 text-xs md:text-sm leading-6 ${bg}`}
        >
          <div className="relative ml-2 md:ml-4 w-7 h-7 text-white flex items-center justify-center">
            <Icon className="w-5 h-5 md:w-6 md:h-6" />
          </div>
          <span className="flex-1 text-base font-medium text-white dark:text-slate-200">
            {domToReact(admTitle.children, options)}
          </span>
          {trigger == "!" && (
            <div
              className="absolute top-2 right-0 md:h-7 flex items-center md:pr-4 pr-2"
              onClick={toggleBodyVisibility}
            >
              <LuChevronLeftSquare
                className={`w-5 h-5 md:w-6 md:h-6 cursor-pointer text-slate-100 dark:text-slate-300 transform ${
                  isBodyVisible ? "" : "rotate-270"
                }`}
              />
            </div>
          )}
        </div>
        <motion.div
          initial={{ opacity: 0, height: 0 }}
          animate={{
            opacity: isBodyVisible ? 1 : 0,
            height: isBodyVisible ? "auto" : 0,
          }}
          exit={{ opacity: 0, height: 0 }}
          layout // Enables smooth transitions between dynamic height changes
          transition={{ duration: 0.5, ease: "easeInOut" }}
          className="px-4 md:px-6 overflow-hidden scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
scrollbar-thumb-slate-200 scrollbar-track-slate-100
dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
        >
          {domToReact(admBody.children, options)}
        </motion.div>
      </div>
    );
  } else {
    return <></>;
  }
}
