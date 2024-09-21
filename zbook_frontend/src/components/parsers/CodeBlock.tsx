"use client";
import { MdOutlineContentCopy } from "react-icons/md";
import { AiFillCode } from "react-icons/ai";
import CodeMermaid from "./CodeMermaid";
import { toast } from "react-toastify";
import { useTranslations } from "next-intl";
import CodeHighLight from "./CodeHighLight";
import { useCallback } from "react";

export default function CodeBlock({
  lang,
  codeString,
}: {
  lang: string;
  codeString: string;
}) {
  const t = useTranslations("Toast");

  const copy2ClipBoard = useCallback(() => {
    navigator.clipboard.writeText(codeString).then(
      () => {
        toast(t("CopiedClipboard"), {
          type: "success",
          autoClose: 500,
        });
      },
      () => {
        toast(t("CopyClipboardFailed"), {
          type: "error",
          autoClose: 1500,
        });
      }
    );
  }, [codeString, t]);

  return (
    <div className="relative z-10 my-[1.25em] md:pb-4 pb-2 col-span-3 font-base rounded-md dark:bg-slate-800/50 dark:ring-1 ring-slate-200/50 dark:ring-slate-900/10 border-[0.01rem] border-slate-300 dark:border-0">
      <div className="relative py-1 md:py-2 space-x-4 rounded-t-md flex items-center justify-center text-slate-400 text-xs md:text-sm leading-6 border-b-[0.01rem] border-slate-300 dark:border-slate-700/30 dark:bg-slate-800/50">
        <div className="relative ml-2 md:ml-4 w-7 h-7 rounded-full text-white flex items-center justify-center">
          <AiFillCode className="w-5 h-5 md:w-6 md:h-6 text-slate-500 dark:text-slate-200" />
        </div>
        <span className="flex-1 text-base font-medium text-slate-900 dark:text-slate-200">
          {lang}
        </span>
        <div className="absolute top-2 right-0 md:h-7 flex items-center md:pr-4 pr-2">
          <MdOutlineContentCopy
            className="w-5 h-5 md:w-6 md:h-6 cursor-pointer text-slate-500 dark:text-slate-300"
            onClick={copy2ClipBoard}
          />
        </div>
      </div>
      <div
        className="highlight px-2 md:px-4 mt-2 md:mt-4 text-xs md:text-sm text-slate-800 dark:text-slate-200 overflow-x-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
      >
        {lang === "mermaid" ? (
          <CodeMermaid graphDefinition={codeString} />
        ) : (
          <CodeHighLight lang={lang} codeString={codeString} />
        )}
      </div>
    </div>
  );
}
