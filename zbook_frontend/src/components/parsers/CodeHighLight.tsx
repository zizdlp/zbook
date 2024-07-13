"use client";
import { useTheme } from "next-themes";
import SyntaxHighlighter from "react-syntax-highlighter";
import {
  atomOneLight,
  atomOneDark,
} from "react-syntax-highlighter/dist/esm/styles/hljs";
import { useEffect, useState, useMemo } from "react";
import { MdOutlineCode } from "react-icons/md";

interface CodeHighLightProps {
  codeString: string;
  lang: string;
}

function truncateLastNewline(str: string) {
  return str.endsWith("\n") ? str.slice(0, -1) : str;
}

export default function CodeHighLight({ codeString, lang }: CodeHighLightProps) {
  const { theme } = useTheme();
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  const style = useMemo(
    () => (theme === "dark" ? atomOneDark : atomOneLight),
    [theme]
  );

  if (!mounted) {
    return (
      <MdOutlineCode className="w-full h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
    );
  }

  return (
    <SyntaxHighlighter
      language={lang}
      showLineNumbers={true}
      wrapLines={true}
      style={style}
      customStyle={{ background: "none", padding: "0" }}
    >
      {truncateLastNewline(codeString)}
    </SyntaxHighlighter>
  );
}