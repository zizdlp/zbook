"use client";

import MathDisplay from "./MathDisplay";
import MathInline from "./MathInline";
import CodeBlock from "./CodeBlock";
import ImageWithFallback from "./ImageWithFallback";
import { ThemeColor } from "../TableOfContent";
import VideoWithFallBack from "./VideoWithFallBack";
import MarkdownImageClient from "./MarkdownImageClient";

interface HtmlParserProps {
  htmlString: string;
  prefixPath: string;
  username: string;
  repo_name: string;
  theme_color: ThemeColor;
  agent: string;
}
function isSafari(userAgent: string): boolean {
  const isSafari = /^((?!chrome|android|crios|fxios).)*safari/i.test(userAgent);
  return isSafari;
}
import { StrictMode } from "react";
import parse, {
  Element,
  domToReact,
  attributesToProps,
} from "html-react-parser";
import type { DOMNode, HTMLReactParserOptions } from "html-react-parser";
import Admonition from "./Admonition";
function getHeadColorClasses(color: ThemeColor) {
  return {
    activeClass: `text-${color}-700 dark:text-${color}-400`,
  };
}
const HtmlParser: React.FC<HtmlParserProps> = ({
  htmlString,
  prefixPath,
  username,
  repo_name,
  theme_color,
  agent,
}) => {
  const options: HTMLReactParserOptions = {
    replace(domNode: DOMNode) {
      if (domNode instanceof Element && domNode.attribs) {
        const tagName = domNode.tagName.toLowerCase();
        const props = attributesToProps(domNode.attribs);
        const randomKey = Math.random().toString(36).substring(2);
        // Check for the class attribute in the attributes array
        const classAttr = domNode.attributes.find(
          (attr) => attr.name === "class"
        );
        const classList = classAttr ? classAttr.value.split(" ") : [];

        if (/^h\d$/.test(tagName)) {
          const level = parseInt(tagName[1], 10);
          const HeadingComponent = `h${level}` as keyof JSX.IntrinsicElements;
          let { activeClass } = getHeadColorClasses(theme_color);
          if (level === 1) {
            return (
              <div key={randomKey}>
                <div
                  className={`h-5 ${activeClass} text-sm font-semibold pb-4`}
                >
                  {prefixPath}
                </div>
                <h1 className="text-2xl sm:text-3xl font-extrabold">
                  {domToReact(domNode.children, options)}
                </h1>
              </div>
            );
          }
          return (
            <HeadingComponent {...props}>
              {domToReact(domNode.children, options)}
            </HeadingComponent>
          );
        } else if (
          classList.includes("math") &&
          classList.includes("display")
        ) {
          return (
            <MathDisplay key={randomKey}>
              {domToReact(domNode.children, options)}
            </MathDisplay>
          );
        } else if (classList.includes("math") && classList.includes("inline")) {
          return (
            <MathInline key={randomKey}>
              {domToReact(domNode.children, options)}
            </MathInline>
          );
        } else if (tagName === "code") {
          const parent = domNode.parentNode as Element | null; // Get the parent node
          const parentIsPre = parent && parent.tagName.toLowerCase() === "pre"; // Check if the parent is a <pre> element

          // If parent is a <pre>, handle as a CodeBlock
          if (parentIsPre) {
            const classAttr = domNode.attributes.find(
              (attr) => attr.name === "class"
            );
            const lang = classAttr
              ? classAttr.value.substring("language-".length)
              : "";

            const codeString = Array.from(domNode.childNodes)
              .map((child) => {
                return (child as unknown as Text).data; // Use TypeScript's type assertion for Text
              })
              .join("");

            return (
              <CodeBlock key={randomKey} codeString={codeString} lang={lang} />
            );
          } else {
            // Otherwise, handle as inline code
            return (
              <code key={randomKey} className="font-jetbrains" {...props}>
                {domToReact(domNode.children, options)}
              </code>
            );
          }
        } else if (tagName === "div") {
          if (classList.includes("admonition")) {
            return <Admonition domNode={domNode} options={options} />;
          }
        } else if (tagName === "p") {
          return (
            <p
              key={randomKey}
              className="overflow-x-auto overflow-y-hidden scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
            scrollbar-thumb-slate-200 scrollbar-track-slate-100
            dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
            >
              {domToReact(domNode.children, options)}
            </p>
          );
        } else if (tagName === "img") {
          const node = domNode as Element;
          const srcAttribute = node.attribs.src; // Access attributes using `attribs`
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
              <MarkdownImageClient
                path={`${prefixPath}/${srcAttribute}`}
                username={username}
                repo_name={repo_name}
                key={randomKey} // Add key for the component
              />
            );
          }
        } else if (tagName === "a") {
          const node = domNode as Element;
          const parent = node.parentNode as Element; // Cast parent to Element
          const urlhref = node.attribs.href; // Access href attribute

          // Process the URL
          let processedUrl = urlhref;
          if (
            processedUrl &&
            (processedUrl.startsWith("http") || processedUrl.startsWith("#"))
          ) {
            // URL is valid
          } else if (processedUrl && processedUrl.endsWith(".md")) {
            processedUrl = `/workspace/${username}/o/${repo_name}/${prefixPath}/${processedUrl.substring(0, processedUrl.length - 3)}`;
          }

          // // Check parent tag
          if (parent && parent.tagName.toUpperCase() === "SUP") {
            return (
              <a
                key={randomKey}
                className={`text-${theme_color}-500 hover:text-${theme_color}-600 text-xs no-underline`}
                href={processedUrl || ""}
              >
                {`[${domToReact(domNode.children, options)}]`}
              </a>
            );
          }

          return (
            <a
              key={randomKey}
              className="underline hover:decoration-2 underline-offset-[0.22rem] text-slate-700 dark:text-gray-200"
              href={processedUrl || ""}
            >
              {domToReact(domNode.children, options)}
            </a>
          );
        } else if (tagName === "iframe") {
          return <VideoWithFallBack src={props.src} key={randomKey} />;
        }
      }

      if (
        domNode instanceof Element &&
        domNode.attribs &&
        domNode.name === "table"
      ) {
        const props = attributesToProps(domNode.attribs);
        return (
          <div className="mt-4 -mb-3">
            <div className="relative bg-slate-50 rounded-xl overflow-hidden dark:bg-slate-800/25">
              <div
                className="absolute inset-0 bg-grid-light dark:bg-grid-dark [mask-image:linear-gradient(0deg,rgba(255,255,255,0.1),rgba(255,255,255,0.6))]  dark:[mask-image:linear-gradient(0deg,rgba(255,255,255,0.1),rgba(255,255,255,0.5))]"
                style={{ backgroundPosition: "10px 10px" }} // 使用对象格式
              ></div>
              <div
                className="relative rounded-xl overflow-x-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
                scrollbar-thumb-slate-200 scrollbar-track-slate-100
                dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
              >
                <div className="shadow-sm my-8">
                  <table
                    className="border-collapse table-auto w-full text-sm"
                    {...props}
                  >
                    {domToReact(domNode.children, options)}
                  </table>
                </div>
              </div>
              <div className="absolute inset-0 pointer-events-none border border-black/10 rounded-xl dark:border-white/5"></div>
            </div>
          </div>
        );
      } else if (
        domNode instanceof Element &&
        domNode.attribs &&
        domNode.name === "thead"
      ) {
        const props = attributesToProps(domNode.attribs);
        return (
          <thead {...props}>{domToReact(domNode.children, options)}</thead>
        );
      } else if (
        domNode instanceof Element &&
        domNode.attribs &&
        domNode.name === "th"
      ) {
        const props = attributesToProps(domNode.attribs);
        const parent = domNode.parent;

        if (parent && parent instanceof Element && parent.name === "tr") {
          const siblings = parent.children.filter(
            (child) => child instanceof Element && child.name === "th"
          );
          const index = siblings.indexOf(domNode);

          let className =
            "border-b dark:border-slate-600 font-medium p-4 pt-0 pb-3   text-left";
          if (index === 0) {
            className += " pl-8";
          } else if (index === siblings.length - 1) {
            className += " pr-8";
          }
          return (
            <th className={className} {...props}>
              {domToReact(domNode.children, options)}
            </th>
          );
        }
        return (
          <th
            className="border-b dark:border-slate-600 font-medium p-4 pt-0 pb-3 text-left"
            {...props}
          >
            {domToReact(domNode.children, options)}
          </th>
        );
      } else if (
        domNode instanceof Element &&
        domNode.attribs &&
        domNode.name === "td"
      ) {
        const props = attributesToProps(domNode.attribs);
        const parent = domNode.parent;

        if (parent && parent instanceof Element && parent.name === "tr") {
          const siblings = parent.children.filter(
            (child) => child instanceof Element && child.name === "td"
          );
          const index = siblings.indexOf(domNode);

          let className =
            "border-b border-slate-200 dark:border-slate-700 p-4  text-left font-normal text-sm";

          if (index === 0) {
            className += " pl-8";
          } else if (index === siblings.length - 1) {
            className += " pr-8";
          }
          return (
            <th className={className} {...props}>
              {domToReact(domNode.children, options)}
            </th>
          );
        }
        return (
          <td
            className="border-b border-slate-100 dark:border-slate-700 p-4 text-left font-normal text-sm"
            {...props}
          >
            {domToReact(domNode.children, options)}
          </td>
        );
      } else if (
        domNode instanceof Element &&
        domNode.attribs &&
        domNode.name === "tbody"
      ) {
        const props = attributesToProps(domNode.attribs);
        return (
          <tbody className="bg-white dark:bg-slate-800" {...props}>
            {domToReact(domNode.children, options)}
          </tbody>
        );
      }
    },
  };

  return (
    <div
      className={`prose prose-sm lg:prose-base prose-slate max-w-none dark:prose-invert dark:text-slate-400
    prose-th:lg:ps-8 prose-th:ps-4 ${isSafari(agent) ? "prose-ol:list-inside" : ""}
    prose-pre:bg-inherit prose-pre:m-0 prose-pre:p-0 prose-table:p-0 prose-table:m-0`}
    >
      <StrictMode>{parse(htmlString, options)}</StrictMode>
    </div>
  );
};

export default HtmlParser;
