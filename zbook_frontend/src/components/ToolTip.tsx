"use client";
import Tippy from "@tippyjs/react";
import "tippy.js/dist/tippy.css";
import "tippy.js/themes/material.css";
import "tippy.js/animations/scale.css";
import React, { ReactElement } from "react";

export default function ToolTip({
  message,
  children,
}: {
  message: string;
  children: ReactElement;
}) {
  return (
    <Tippy content={message} theme={"material"} animation={"scale"}>
      {children}
    </Tippy>
  );
}
