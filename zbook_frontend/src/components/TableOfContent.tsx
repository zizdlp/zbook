"use client";
import parse from "html-react-parser";

import React, { useState, useEffect } from "react";

import {
  domToReact,
  attributesToProps,
  DOMNode,
  Element,
  HTMLReactParserOptions,
} from "html-react-parser";
import ContentSideBar from "./sidebars/ContentSideBar";
interface TableOfContentProps {
  sectionIds: string[];
  markdownlist: string;
}

export default function TableOfContent(props: TableOfContentProps) {
  const [activeSectionId, setActiveSectionId] = useState("");
  function get_act() {
    return activeSectionId;
  }
  const handleClick = (e: React.ChangeEvent<any>) => {
    e.preventDefault();
    var href = e.target.getAttribute("href");
    const escapedSelector = href.replace(/#(\d)/, "#\\3$1 ");
    const section1 = document.querySelector(escapedSelector);
    const navHeight = 60;
    window.scrollTo({
      top: section1.offsetTop - navHeight,
      behavior: "smooth",
    });
  };
  useEffect(() => {
    const handleScroll = () => {
      const visibleSectionId = props.sectionIds.find((sectionId) => {
        const section = document.getElementById(sectionId);
        if (section) {
          const rect = section.getBoundingClientRect();
          return rect.top >= 70 && rect.top <= 100;
        }
        return false;
      });
      if (visibleSectionId != undefined) {
        setActiveSectionId(visibleSectionId);
      }
    };

    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, [props.sectionIds]);
  var html_parser_options_list: HTMLReactParserOptions = {
    replace: (domNode: DOMNode) => {
      if (domNode instanceof Element && domNode.name === "ul") {
        // 访问父元素的属性
        const parent = domNode.parent;
        //@ts-ignore
        const parentAttr = parent != undefined && parent.attribs != undefined;
        if (parentAttr) {
          return (
            <ul className="text-slate-700 leading-6">
              {domToReact(domNode.children, html_parser_options_list)}
            </ul>
          );
        } else {
          return (
            <ul className="text-slate-700 border-l text-sm font-normal border-slate-100 dark:border-slate-700">
              {domToReact(domNode.children, html_parser_options_list)}
            </ul>
          );
        }
      }
      if (domNode instanceof Element && domNode.name === "li") {
        // 访问父元素的属性
        const parent = domNode.parent;
        //@ts-ignore
        const parentAttr = parent != undefined && parent.attribs != undefined;

        if (parentAttr) {
          //@ts-ignore
          const id = parent.attribs.id;
          if (id && id == "content_title") {
            if (domNode.firstChild?.nextSibling != undefined) {
              let title = domNode.firstChild?.nextSibling as Element;
              if (title?.name == "a") {
                return (
                  <li>
                    {domToReact(
                      domNode.children.slice(2, -1),
                      html_parser_options_list
                    )}
                  </li>
                );
              }
            }
          }
        }
        return (
          <li className="ml-4">
            {domToReact(domNode.children, html_parser_options_list)}
          </li>
        );
      } else if (domNode instanceof Element && domNode.name === "a") {
        const props = attributesToProps(domNode.attribs);
        // 访问父元素的属性
        const parent = domNode.parent;
        if (!parent || !parent.parent || !parent.parent.parent) {
          return (
            <a
              {...props}
              onClick={handleClick}
              className={`group flex items-start py-1 hover:text-sky-500  dark:hover:text-sky-500 overflow-scroll ${
                get_act() === props.href.substring(1)
                  ? "dark:text-sky-500 text-sky-500"
                  : "dark:text-slate-400"
              }`}
            >
              {domToReact(domNode.children, html_parser_options_list)}
            </a>
          );
        }
        return (
          <a
            {...props}
            onClick={handleClick}
            className={`group flex items-start py-1 hover:text-sky-500  dark:hover:text-sky-500 overflow-scroll ${
              get_act() === props.href.substring(1)
                ? "dark:text-sky-500 text-sky-500"
                : "dark:text-slate-400"
            }`}
          >
            {domToReact(domNode.children, html_parser_options_list)}
          </a>
        );
      }
    },
  };

  return (
    <ContentSideBar>
      {parse(props.markdownlist, html_parser_options_list)}
    </ContentSideBar>
  );
}
