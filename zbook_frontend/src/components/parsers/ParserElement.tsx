import { StrictMode } from "react";
import parse, {
  Element,
  domToReact,
  attributesToProps,
} from "html-react-parser";
import type { DOMNode, HTMLReactParserOptions } from "html-react-parser";

const options: HTMLReactParserOptions = {
  replace(domNode: DOMNode) {
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
            <div className="relative rounded-xl overflow-auto">
              <div className="shadow-sm my-8 overflow-scroll">
                <table
                  className="border-collapse table-auto w-full text-sm"
                  {...props}
                >
                  {domToReact(domNode.children, options)}
                </table>
              </div>
            </div>
            <div className="absolute inset-0 pointer-events-none border border-black/5 rounded-xl dark:border-white/5"></div>
          </div>
        </div>
      );
    } else if (
      domNode instanceof Element &&
      domNode.attribs &&
      domNode.name === "thead"
    ) {
      const props = attributesToProps(domNode.attribs);
      return <thead {...props}>{domToReact(domNode.children, options)}</thead>;
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
      domNode.name === "p"
    ) {
      const props = attributesToProps(domNode.attribs);
      return (
        <p className="not-prose" {...props}>
          {domToReact(domNode.children, options)}
        </p>
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
          "border-b border-slate-100 dark:border-slate-700 p-4  text-left font-normal text-sm";

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
      domNode.name === "tr"
    ) {
      const props = attributesToProps(domNode.attribs);
      return <tr {...props}>{domToReact(domNode.children, options)}</tr>;
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
const ParserElement = ({ node }: { node: any }) => {
  return <StrictMode>{parse(node.outerHTML, options)}</StrictMode>;
};

export default ParserElement;
