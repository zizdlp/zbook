"use client";

import React, { createContext, useState } from "react";
export const SideBarContext = createContext<{
  sideBarOpen: boolean;
  setSideBarOpen: React.Dispatch<React.SetStateAction<boolean>>;
}>({
  sideBarOpen: false,
  setSideBarOpen: () => {},
});

export default function SideBarProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [sideBarOpen, setSideBarOpen] = useState(false);

  // Context values passed to consumer
  const value = {
    sideBarOpen, // <------ Expose Value to Consumer
    setSideBarOpen, // <------ Expose Setter to Consumer
  };
  return (
    <SideBarContext.Provider value={value}>{children}</SideBarContext.Provider>
  );
}
