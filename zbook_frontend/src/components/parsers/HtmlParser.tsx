/* eslint-disable react/jsx-no-literals */
import { JSDOM } from "jsdom";
import React, { Suspense } from "react";
import MathDisplay from "./MathDisplay";
import MathInline from "./MathInline";
import MarkdownImage from "@/components/parsers/MarkdownImage";
import { BsFillBookmarkCheckFill } from "react-icons/bs";
import { TiWarning } from "react-icons/ti";
import { MdError } from "react-icons/md";
import { FaInfoCircle } from "react-icons/fa";
import { MdTipsAndUpdates } from "react-icons/md";
import { CiImageOn } from "react-icons/ci";
import ParserElement from "./ParserElement";
import CodeBlock from "./CodeBlock";
import ImageWithFallback from "./ImageWithFallback";
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
const parseHTMLString = (
  htmlString: string,
  prefixPath: string,
  username: string,
  repo_name: string
): React.ReactNode => {
  const { window } = new JSDOM("");
  const parser = new window.DOMParser();
  const doc = parser.parseFromString(htmlString, "text/html");
  const processNode = (node: Node): React.ReactNode => {
    if (node instanceof window.Element) {
      const tagName = node.tagName.toUpperCase();
      const idAttribute = node.getAttribute("id");
      const randomKey = Math.random().toString(36).substring(2);
      if (tagName.startsWith("H") && !isNaN(parseInt(tagName[1], 10))) {
        const HeadingComponent =
          tagName.toLowerCase() as keyof JSX.IntrinsicElements;
        const level = parseInt(tagName.substring(1), 10);
        if (level == 1) {
          return (
            <div key={randomKey}>
              <div className="h-5 text-purple-700 dark:text-purple-400 text-sm font-semibold">
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
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <span key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </span>
        );
      } else if (tagName === "CODE") {
        const props = attributesToProps(node.attributes);
        const randomKey = Math.random().toString(36).substring(2);
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
        const props = attributesToProps(node.attributes);
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <pre key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </pre>
        );
      } else if (tagName === "HR") {
        const randomKey = Math.random().toString(36).substring(2);
        return <hr key={randomKey} />;
      } else if (tagName === "UL") {
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <ul key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </ul>
        );
      } else if (tagName === "OL") {
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <ol key={randomKey}>
            {Array.from(node.childNodes).map(processNode)}
          </ol>
        );
      } else if (tagName === "LI") {
        const props = attributesToProps(node.attributes);
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <li key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </li>
        );
      } else if (tagName === "SUP") {
        const props = attributesToProps(node.attributes);
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <sup key={randomKey} {...props}>
            {Array.from(node.childNodes).map(processNode)}
          </sup>
        );
      }

      if (tagName === "DIV") {
        const classAttribute = node.getAttribute("class");

        if (classAttribute === "footnotes") {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <div className="" key={randomKey}>
              <ParserElement node={node} key={randomKey} />
            </div>
          );
        } else if (classAttribute === "adm-title") {
          type ParentType = "note" | "warning" | "info" | "tip" | "error";
          let parentType: ParentType = "note"; // or "warning", "info", "tip", "error"

          const iconTypes = {
            note: BsFillBookmarkCheckFill,
            warning: TiWarning,
            info: FaInfoCircle,
            tip: MdTipsAndUpdates,
            error: MdError,
          };

          let bg1 = "bg-cyan-600/75 dark:bg-cyan-500/25";
          const parent = node.parentElement;
          const parentAttr =
            parent !== null && parent.getAttribute !== undefined;

          if (parentAttr) {
            const attrsArray = Array.from(parent.attributes);
            const classAttribute = attrsArray.find(
              (attr) => attr.name === "class"
            );

            if (
              classAttribute &&
              classAttribute.value === "admonition adm-note"
            ) {
              parentType = "note";
              bg1 = "bg-cyan-600/75 dark:bg-cyan-500/25";
            } else if (
              classAttribute &&
              classAttribute.value === "admonition adm-warning"
            ) {
              parentType = "warning";
              bg1 = "bg-yellow-600/75 dark:bg-yellow-500/25";
            } else if (
              classAttribute &&
              classAttribute.value === "admonition adm-info"
            ) {
              parentType = "info";
              bg1 = "bg-green-600/75 dark:bg-green-500/25";
            } else if (
              classAttribute &&
              classAttribute.value === "admonition adm-tip"
            ) {
              parentType = "tip";
              bg1 = "bg-lime-600/75 dark:bg-lime-500/25";
            } else if (
              classAttribute &&
              classAttribute.value === "admonition adm-error"
            ) {
              parentType = "error";
              bg1 = "bg-red-600/75 dark:bg-red-500/25";
            }
          }

          const Icon = iconTypes[parentType];
          const randomKey = Math.random().toString(36).substring(2);
          const randomKey2 = Math.random().toString(36).substring(2);
          return (
            <div
              key={randomKey}
              className={`relative py-1 md:py-2 space-x-4 rounded-t-md flex items-center justify-center text-slate-400 text-xs md:text-sm leading-6  ${bg1}`}
            >
              <div
                key={randomKey}
                className={`relative ml-2 md:ml-4 w-7 h-7  text-white flex items-center justify-center `}
              >
                <Icon className="w-5 h-5 md:w-6 md:h-6" />
              </div>
              <span
                key={randomKey2}
                className="flex-1 text-base font-medium text-white dark:text-slate-200"
              >
                {Array.from(node.childNodes).map(processNode)}
              </span>
            </div>
          );
        } else if (classAttribute === "adm-body") {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <div key={randomKey} className="px-6 overflow-x-auto">
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        } else if (classAttribute && classAttribute.includes("admonition")) {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <div
              key={randomKey}
              className="my-[2em] rounded-md bg-slate-100/50 dark:bg-slate-800/50 ring-1 ring-slate-200/50 dark:ring-slate-900/10"
            >
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        } else {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <div key={randomKey} className="text-xs md:text-base">
              {Array.from(node.childNodes).map(processNode)}
            </div>
          );
        }
      } else if (tagName === "P") {
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <p key={randomKey}>{Array.from(node.childNodes).map(processNode)}</p>
        );
      } else if (tagName === "TABLE") {
        const randomKey = Math.random().toString(36).substring(2);
        return <ParserElement key={randomKey} node={node} />;
      } else if (tagName === "BLOCKQUOTE") {
        const randomKey = Math.random().toString(36).substring(2);
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
                <CiImageOn className="w-full h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
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
        const randomKey = Math.random().toString(36).substring(2);
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
          urlhref = "/workspace/" + username + "/o/" + repo_name;
          "/" + prefixPath + "/" + urlhref.substring(0, urlhref.length - 3);
        }
        if (parentAttr) {
          const tag = parent.tagName;
          const randomKey = Math.random().toString(36).substring(2);
          if (tag === "SUP") {
            return (
              <a
                key={randomKey}
                className="text-purple-500 hover:text-purple-600 text-xs"
                href={urlhref || ""}
              >
                [{Array.from(node.childNodes).map(processNode)}]
              </a>
            );
          }
        }
        const href = node.getAttribute("href");
        if (href && href.includes("footnote")) {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <a
              key={randomKey}
              className="text-red-500 hover:text-red-600"
              href={urlhref || ""}
            >
              {Array.from(node.childNodes).map(processNode)}
            </a>
          );
        } else {
          const randomKey = Math.random().toString(36).substring(2);
          return (
            <a
              key={randomKey}
              className="px-1 underline font-semibold decoration-purple-500 dark:decoration-purple-400 hover:decoration-2 underline-offset-[0.22rem] text-slate-700 dark:text-gray-200"
              href={urlhref || ""}
            >
              {Array.from(node.childNodes).map(processNode)}
            </a>
          );
        }
      } else if (tagName === "STRONG") {
        const randomKey = Math.random().toString(36).substring(2);
        return (
          <strong key={randomKey} className="font-semibold dark:text-gray-200">
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
}

const HtmlParser: React.FC<HtmlParserProps> = ({
  htmlString,
  prefixPath,
  username,
  repo_name,
}) => {
  const parsedHTML = parseHTMLString(
    htmlString,
    prefixPath,
    username,
    repo_name
  );
  return (
    <div
      className="prose prose-sm lg:prose-base prose-slate max-w-none dark:prose-invert dark:text-slate-400
    prose-pre:bg-inherit prose-pre:m-0 prose-pre:p-0 prose-table:p-0 prose-table:m-0"
    >
      {parsedHTML}
    </div>
  );
};

export default HtmlParser;
