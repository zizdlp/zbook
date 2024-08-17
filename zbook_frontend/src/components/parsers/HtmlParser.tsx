/* eslint-disable react/jsx-no-literals */
import { JSDOM } from "jsdom";
import React, { Suspense } from "react";
import MathDisplay from "./MathDisplay";
import MathInline from "./MathInline";
import MarkdownImage from "@/components/parsers/MarkdownImage";
import { CiImageOn } from "react-icons/ci";
import ParserElement from "./ParserElement";
import CodeBlock from "./CodeBlock";
import ImageWithFallback from "./ImageWithFallback";
import { BsFillBookmarkCheckFill } from "react-icons/bs";
import { FaInfoCircle } from "react-icons/fa";
import { MdError, MdTipsAndUpdates } from "react-icons/md";
import { TiWarning } from "react-icons/ti";
import { ThemeColor } from "../TableOfContent";

interface Attribute {
  name: string;
  value: string;
}

function attributesToProps(attributes: NamedNodeMap): {
  [key: string]: string;
} {
  const props: { [key: string]: string } = {};

  for (let i = 0; i < attributes.length; i++) {
    const attribute: Attribute = attributes[i] as Attribute;
    props[attribute.name] = attribute.value;
  }

  return props;
}
function getHeadColorClasses(color: ThemeColor) {
  return {
    activeClass: `text-${color}-700 dark:text-${color}-400`,
  };
}
const parseHTMLString = (
  htmlString: string,
  prefixPath: string,
  username: string,
  repo_name: string,
  theme_color: ThemeColor
): React.ReactNode => {
  const { window } = new JSDOM("");
  const parser = new window.DOMParser();
  const doc = parser.parseFromString(htmlString, "text/html");
  let { activeClass } = getHeadColorClasses(theme_color);
  const processNode = (node: Node): React.ReactNode => {
    if (node instanceof window.Element) {
      const tagName = node.tagName.toUpperCase();
      const idAttribute = node.getAttribute("id");
      const randomKey = Math.random().toString(36).substring(2);
      const props = attributesToProps(node.attributes);
      if (tagName.startsWith("H") && !isNaN(parseInt(tagName[1], 10))) {
        const HeadingComponent =
          tagName.toLowerCase() as keyof JSX.IntrinsicElements;
        const level = parseInt(tagName.substring(1), 10);
        if (level == 1) {
          return (
            <div key={randomKey}>
              <div className={`h-5 ${activeClass} text-sm font-semibold`}>
                {prefixPath}
              </div>
              <h1 className="text-2xl sm:text-3xl font-extrabold">
                {Array.from(node.childNodes).map(processNode)}
              </h1>
            </div>
          );
        }

        return (
          <HeadingComponent key={randomKey} id={idAttribute || undefined}>
            {Array.from(node.childNodes).map(processNode)}
          </HeadingComponent>
        );
      } else if (
        node.classList.contains("math") &&
        node.classList.contains("display")
      ) {
        return (
          <MathDisplay key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </MathDisplay>
        );
      } else if (
        node.classList.contains("math") &&
        node.classList.contains("inline")
      ) {
        return (
          <MathInline key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </MathInline>
        );
      } else if (tagName === "SPAN") {
        return (
          <span key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </span>
        );
      } else if (tagName === "CODE") {
        const parent = node.parentElement;
        const parentIsPre = parent?.tagName.toUpperCase() === "PRE";
        if (parentIsPre) {
          const lang = node.classList.value.substring("language-".length);
          return (
            <CodeBlock
              key={randomKey}
              codeString={node.textContent ?? ""}
              lang={lang}
            />
          );
        } else {
          return (
            <code key={randomKey} className="font-jetbrains" {...props}>
              {Array.from(node.childNodes).map(processNode)}
            </code>
          );
        }
      } else if (tagName === "PRE") {
        return (
          <pre key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </pre>
        );
      } else if (tagName === "HR") {
        return <hr key={randomKey} />;
      } else if (tagName === "UL") {
        return (
          <ul key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </ul>
        );
      } else if (tagName === "OL") {
        return (
          <ol key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </ol>
        );
      } else if (tagName === "LI") {
        return (
          <li key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </li>
        );
      } else if (tagName === "SUP") {
        return (
          <sup key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </sup>
        );
      }

      if (tagName === "DIV") {
        const classAttribute = node.getAttribute("class");
        if (classAttribute === "adm-title") {
          type ParentType = "note" | "warning" | "info" | "tip" | "error";

          const iconTypes = {
            note: BsFillBookmarkCheckFill,
            warning: TiWarning,
            info: FaInfoCircle,
            tip: MdTipsAndUpdates,
            error: MdError,
          };

          const bgColors: Record<ParentType, string> = {
            note: "bg-cyan-600/75 dark:bg-cyan-500/25",
            warning: "bg-yellow-600/75 dark:bg-yellow-500/25",
            info: "bg-green-600/75 dark:bg-green-500/25",
            tip: "bg-lime-600/75 dark:bg-lime-500/25",
            error: "bg-red-600/75 dark:bg-red-500/25",
          };

          const getParentTypeAndBg = (
            parent: HTMLElement | null
          ): { parentType: ParentType; bg: string } => {
            if (!parent) return { parentType: "note", bg: bgColors.note };
            const classList = parent.getAttribute("class");
            if (classList) {
              if (classList.includes("adm-note"))
                return { parentType: "note", bg: bgColors.note };
              if (classList.includes("adm-warning"))
                return { parentType: "warning", bg: bgColors.warning };
              if (classList.includes("adm-info"))
                return { parentType: "info", bg: bgColors.info };
              if (classList.includes("adm-tip"))
                return { parentType: "tip", bg: bgColors.tip };
              if (classList.includes("adm-error"))
                return { parentType: "error", bg: bgColors.error };
            }

            return { parentType: "note", bg: bgColors.note };
          };
          const { parentType, bg } = getParentTypeAndBg(node.parentElement);
          const Icon = iconTypes[parentType];
          return (
            <div
              key={randomKey}
              className={`relative py-1 md:py-2 space-x-4 rounded-t-md flex items-center justify-center text-slate-400 text-xs md:text-sm leading-6 ${bg}`}
            >
              <div className="relative ml-2 md:ml-4 w-7 h-7 text-white flex items-center justify-center">
                <Icon className="w-5 h-5 md:w-6 md:h-6" />
              </div>
              <span className="flex-1 text-base font-medium text-white dark:text-slate-200">
                {Array.from(node.childNodes).map(processNode)}
              </span>
            </div>
          );
        } else if (classAttribute === "adm-body") {
          return (
            <div key={randomKey} className="px-4 md:px-6 overflow-x-auto">
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        } else if (classAttribute && classAttribute.includes("admonition")) {
          return (
            <div
              key={randomKey}
              className="my-[1.25em] bg-slate-100/50 dark:bg-slate-800/50 ring-1 ring-slate-200/50 dark:ring-slate-900/10 rounded-md"
            >
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        } else if (classAttribute === "footnotes") {
          return <ParserElement key={randomKey} node={node} />;
        } else {
          return (
            <div key={randomKey}>
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        }
      } else if (tagName === "P") {
        return (
          <p key={randomKey} className="overflow-scroll">
            {Array.from(node.childNodes).map(processNode)}
          </p>
        );
      } else if (tagName === "TABLE") {
        return <ParserElement key={randomKey} node={node} />;
      } else if (tagName === "BLOCKQUOTE") {
        return (
          <blockquote key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </blockquote>
        );
      } else if (tagName === "IMG") {
        const srcAttribute = (node as Element).getAttribute("src");

        if (srcAttribute && srcAttribute.startsWith("http")) {
          return (
            <ImageWithFallback
              src={srcAttribute}
              alt="Landscape picture"
              key={randomKey}
            />
          );
        } else {
          return (
            <Suspense
              key={randomKey}
              fallback={
                <CiImageOn className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
              }
            >
              <MarkdownImage
                path={prefixPath + "/" + srcAttribute}
                username={username}
                repo_name={repo_name}
              />
            </Suspense>
          );
        }
      } else if (tagName === "DEL") {
        return (
          <del key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </del>
        );
      } else if (tagName === "A") {
        const parent = node.parentElement;
        const parentAttr = parent !== null && parent.getAttribute !== undefined;
        let urlhref = node.getAttribute("href");
        if (
          urlhref &&
          (urlhref.startsWith("http") || urlhref.startsWith("#"))
        ) {
        } else if (urlhref && urlhref.endsWith(".md")) {
          urlhref =
            "/workspace/" +
            username +
            "/o/" +
            repo_name +
            "/" +
            prefixPath +
            "/" +
            urlhref.substring(0, urlhref.length - 3);
        }
        if (parentAttr) {
          const tag = parent.tagName;
          if (tag === "SUP") {
            return (
              <a
                key={randomKey}
                className={`text-${theme_color}-500 hover:text-${theme_color}-600 text-xs no-underline`}
                href={urlhref || ""}
              >
                [{Array.from(node.childNodes).map(processNode)}]
              </a>
            );
          }
        }
        return (
          <a
            key={randomKey}
            className="underline hover:decoration-2 underline-offset-[0.22rem] text-slate-700 dark:text-gray-200"
            href={urlhref || ""}
          >
            {Array.from(node.childNodes).map(processNode)}
          </a>
        );
      } else if (tagName === "STRONG") {
        return (
          <strong key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </strong>
        );
      }
    } else if (node instanceof window.Text) {
      // 处理文本节点
      return node.textContent;
    } else {
      return null; // 或者其他处理逻辑
    }
  };

  // Start processing from the body of the parsed HTML
  return Array.from(doc.body.childNodes).map(processNode);
};

interface HtmlParserProps {
  htmlString: string;
  prefixPath: string;
  username: string;
  repo_name: string;
  theme_color: ThemeColor;
}

const HtmlParser: React.FC<HtmlParserProps> = ({
  htmlString,
  prefixPath,
  username,
  repo_name,
  theme_color,
}) => {
  const parsedHTML = parseHTMLString(
    htmlString,
    prefixPath,
    username,
    repo_name,
    theme_color
  );
  return (
    <div
      className="prose prose-sm lg:prose-base prose-slate max-w-none dark:prose-invert dark:text-slate-400
    prose-th:lg:ps-8 prose-th:ps-4 prose-ol:
    prose-pre:bg-inherit prose-pre:m-0 prose-pre:p-0 prose-table:p-0 prose-table:m-0"
    >
      {parsedHTML}
    </div>
  );
};

export default HtmlParser;
