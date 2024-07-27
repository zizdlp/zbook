"use client";

import { SearchType } from "@/utils/const_value";
import React, { createContext, useState } from "react";

// 更新上下文的类型
export const SearchDialogContext = createContext<{
  searchDialogOpen: boolean;
  setSearchDialogOpen: React.Dispatch<React.SetStateAction<boolean>>;
  searchType: SearchType;
  setSearchType: React.Dispatch<React.SetStateAction<SearchType>>;
}>({
  searchDialogOpen: true,
  setSearchDialogOpen: () => {},
  searchType: SearchType.DOCUMENT,
  setSearchType: () => {},
});

export default function SearchDialogProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [searchDialogOpen, setSearchDialogOpen] = useState(false);
  const [searchType, setSearchType] = useState<SearchType>(SearchType.DOCUMENT);

  // Context values passed to consumer
  const value = {
    searchDialogOpen,
    setSearchDialogOpen,
    searchType,
    setSearchType,
  };

  return (
    <SearchDialogContext.Provider value={value}>
      {children}
    </SearchDialogContext.Provider>
  );
}
