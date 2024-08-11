"use client";
import parse from "html-react-parser";

import React, { useState, useEffect } from "react";
import { LuListMinus } from "react-icons/lu";
export enum ThemeColor {
  Violet = "violet",
  Green = "green",
  Red = "red",
  Yellow = "yellow",
  Teal = "teal",
  Sky = "sky",
  Cyan = "cyan",
  Pink = "pink",
  Indigo = "indigo",
}
function getColorClasses(color: ThemeColor) {
  return {
    hoverClass: `hover:text-${color}-500 hover:dark:text-${color}-400 hover:font-semibold`,
    activeClass: `text-${color}-500 dark:text-${color}-400 font-semibold`,
  };
}
import {
  domToReact,
  attributesToProps,
  DOMNode,
  Element,
  HTMLReactParserOptions,
} from "html-react-parser";
import RightSideBarWrapper from "./sidebars/RightSideBarWrapper";
import { useTranslations } from "next-intl";
interface TableOfContentProps {
  sectionIds: string[];
  markdownlist: string;
  theme_color: ThemeColor;
}

export default function TableOfContent(props: TableOfContentProps) {
  const [activeSectionId, setActiveSectionId] = useState("");
  const t = useTranslations("RightSideBar");
  function get_act() {
    return activeSectionId;
  }
  const handleClick = (e: React.ChangeEvent<any>) => {
    e.preventDefault();
    var href = e.target.getAttribute("href");
    const escapedSelector = href.replace(/#(\d)/, "#\\3$1 ");
    const section1 = document.querySelector(escapedSelector);
    const navHeight = -35;
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
  const { hoverClass, activeClass } = getColorClasses(props.theme_color);
  console.log("hoverClass:", hoverClass);
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
              className={`group flex items-start py-1   ${hoverClass} overflow-scroll ${
                get_act() === props.href.substring(1)
                  ? activeClass
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
            className={`group flex items-start py-1  ${hoverClass} overflow-scroll ${
              get_act() === props.href.substring(1)
                ? activeClass
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
    <RightSideBarWrapper>
      <div
        className="fixed overflow-auto w-[19.5rem]"
        style={{ height: "calc(100vh - 128px)" }}
      >
        {/* 确保父元素具有固定高度或最大高度 */}
        <div className="h-full">
          {/* 子元素充满父元素的高度 */}
          <div className="text-gray-600 dark:text-gray-200 font-medium flex items-center space-x-2">
            <LuListMinus className="h-4 w-4" />
            <span>{t("OnThisPage")}</span>
          </div>
          {parse(props.markdownlist, html_parser_options_list)}
        </div>
      </div>
    </RightSideBarWrapper>
  );
}
