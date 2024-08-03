/* eslint-disable react/jsx-no-literals */
import React from "react";
import { FaBookOpen, FaDiscord, FaGithub } from "react-icons/fa";
import FrpcAnchorItem from "./FrpcAnchorItem";
export default function FrpcAnchor() {
  const menuItems = [
    {
      href: "/introduction",
      text: "Documentation",
      selected: true,
      icon: FaBookOpen,
    },
    {
      href: "https://loopholelabs.io/discord",
      text: "discord",
      selected: false,
      icon: FaDiscord,
    },
    {
      href: "https://github.com/loopholelabs/frpc-go",
      text: "frpc",
      selected: false,
      icon: FaGithub,
    },
  ];

  return (
    <>
      {menuItems.map((item, index) => (
        <FrpcAnchorItem key={index} {...item} />
      ))}
    </>
  );
}
