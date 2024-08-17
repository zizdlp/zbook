"use client";
import Tippy from "@tippyjs/react";
import "tippy.js/dist/tippy.css";
import "tippy.js/themes/material.css";
import "tippy.js/animations/scale.css";
import "tippy.js/themes/light.css";
import React, { ReactElement } from "react";
import { useTheme } from "next-themes";

export default function ToolTip({
  message,
  children,
}: {
  message: string;
  children: ReactElement;
}) {
  const { theme } = useTheme();
  return (
    <Tippy
      content={message}
      theme={theme == "dark" ? "material" : "light"}
      animation={"scale"}
    >
      {children}
    </Tippy>
  );
}
