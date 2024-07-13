"use client";

import React, { createContext, useState } from "react";
export const SearchDialogContext = createContext<{
  searchDialogOpen: boolean;
  setSearchDialogOpen: React.Dispatch<React.SetStateAction<boolean>>;
  searchType: number;
  setSearchType: React.Dispatch<React.SetStateAction<number>>;
}>({
  searchDialogOpen: true,
  setSearchDialogOpen: () => {},
  searchType: 0,
  setSearchType: () => {},
});

export default function SearchDialogProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [searchDialogOpen, setSearchDialogOpen] = useState(false);
  const [searchType, setSearchType] = useState(0);

  // Context values passed to consumer
  const value = {
    searchDialogOpen, // <------ Expose Value to Consumer
    setSearchDialogOpen, // <------ Expose Setter to Consumer
    searchType,
    setSearchType,
  };
  return (
    <SearchDialogContext.Provider value={value}>
      {children}
    </SearchDialogContext.Provider>
  );
}
