"use client";
import mermaid from "mermaid";
import { useContext, useEffect, useState, useCallback } from "react";
import "katex/dist/katex.min.css";
import { useTheme } from "next-themes";
import { OperationContext } from "@/providers/OperationProvider";
import { MdOutlineCode, MdOutlineCodeOff } from "react-icons/md";
import { logger } from "@/utils/logger";

let currentId = 0;
const uuid = () => `mermaid-${(currentId++).toString()}`;

const flowchart = {
  htmlLabels: true,
  curve: "basis",
  useMaxWidth: true,
};

const sequence = {
  diagramMarginX: 50,
  diagramMarginY: 10,
  actorMargin: 50,
  width: 150,
  height: 65,
  boxMargin: 10,
  boxTextMargin: 5,
  noteMargin: 10,
  messageMargin: 35,
  mirrorActors: true,
  bottomMarginAdj: 1,
  useMaxWidth: true,
  rightAngles: false,
  showSequenceNumbers: false,
};

const gantt = {
  titleTopMargin: 25,
  barHeight: 20,
  barGap: 4,
  topPadding: 50,
  leftPadding: 75,
  gridLineStartPadding: 35,
  fontSize: 11,
  fontFamily: '"Open-Sans", "sans-serif"',
  numberSectionStyles: 4,
  axisFormat: "%Y-%m-%d",
};

const DEFAULT_CONFIG = {
  startOnLoad: false,
  logLevel: 4,
  securityLevel: "strict",
  arrowMarkerAbsolute: false,
  flowchart: flowchart,
  sequence: sequence,
  gantt: gantt,
};

const DARK_THEME_CONFIG = {
  ...DEFAULT_CONFIG,
  theme: "dark",
};

const LIGHT_THEME_CONFIG = {
  ...DEFAULT_CONFIG,
  theme: "default",
};

function CodeMermaid({ graphDefinition }: { graphDefinition: string }) {
  const { resolvedTheme } = useTheme();
  const [mounted, setMounted] = useState(false);
  const { mutationToggleTheme } = useContext(OperationContext);
  const [htmlText, setHtmlText] = useState("");

  useEffect(() => {
    setMounted(true);
  }, []);

  const renderMermaid = useCallback(
    async (config: any) => {
      try {
        mermaid.initialize(config);
        const valid = await mermaid.parse(graphDefinition, {
          suppressErrors: true,
        });
        if (valid) {
          const { svg } = await mermaid.render(uuid(), graphDefinition);
          setHtmlText(svg);
        } else {
          setHtmlText("");
        }
      } catch (e) {
        setHtmlText("");
        logger.error(`mermaid render failed:${e}`);
      }
    },
    [graphDefinition]
  );

  useEffect(() => {
    if (mounted && graphDefinition) {
      const config =
        resolvedTheme === "dark" ? DARK_THEME_CONFIG : LIGHT_THEME_CONFIG;
      renderMermaid(config);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mounted, resolvedTheme, mutationToggleTheme]);
  if (!mounted) {
    return (
      <MdOutlineCode className="w-full h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
    );
  }

  if (htmlText) {
    return (
      <div
        className="flex grow items-center justify-center"
        dangerouslySetInnerHTML={{ __html: htmlText }}
      />
    );
  } else {
    return (
      <MdOutlineCodeOff className="rounded-md w-full h-96 py-40 bg-gray-200 dark:bg-gray-700/75 text-slate-500 dark:text-slate-400" />
    );
  }
}

export default CodeMermaid;
