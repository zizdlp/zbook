import { FaDiscord, FaGithub, FaYoutube } from "react-icons/fa";
import { FaBilibili } from "react-icons/fa6";
import { Link } from "@/navigation";
import { getTranslations } from "next-intl/server";
import { beian, doc_reponame, doc_username } from "@/utils/env_variable";
import { IconType } from "react-icons/lib";
import IconItem from "../IconComponent";

const SocialLink = ({
  href,
  ariaLabel,
  Icon,
}: {
  href: string;
  ariaLabel: string;
  Icon: IconType;
}) => (
  <Link className="group" aria-label={ariaLabel} href={href}>
    <Icon className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700" />
  </Link>
);

const FooterSection = ({ title, links }: { title: string; links: any }) => (
  <div className="lg:col-span-2 col-span-4">
    <h2 className="font-semibold text-slate-900 dark:text-slate-100">
      {title}
    </h2>
    <ul className="mt-3 space-y-2">
      {links.map(({ url, label }: { url: string; label: string }) => (
        <li key={url}>
          <Link
            className="hover:text-slate-900 dark:hover:text-slate-300"
            href={url}
          >
            {label}
          </Link>
        </li>
      ))}
    </ul>
  </div>
);

export default async function FooterForHome() {
  const t = await getTranslations("HomePage");

  const sections = [
    {
      title: t("GettingStarted"),
      links: [
        { url: "/auth/login", label: t("Login") },
        { url: "/auth/register", label: t("Register") },
        { url: "/cases", label: t("Cases") },
        {
          url: `/workspace/${doc_username}/o/${doc_reponame}`,
          label: t("Documentation"),
        },
      ],
    },
    {
      title: t("CoreConcepts"),
      links: [
        { url: "#features_section", label: t("FeatureSection") },
        { url: "#dashboard", label: t("DashboardSection") },
        { url: "#create_repo", label: t("CreateRepoSection") },
        { url: "#multi_user_section", label: t("MultiUserSection") },
      ],
    },
    {
      title: t("Community"),
      links: [
        { url: "https://github.com/zizdlp/zbook", label: "GitHub" },
        { url: "https://space.bilibili.com/1448262500", label: "Bilibili" },
        {
          url: "https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg",
          label: "YouTube",
        },
        {
          url: "https://discord.com/channels/1250069935594536960/1250069935594536963",
          label: "Discord",
        },
      ],
    },
    {
      title: t("Links"),
      links: [
        { url: "/terms", label: t("Terms") },
        { url: "/privacy", label: t("Privacy") },
        { url: "https://linchat.zizdlp.com", label: t("LinChatWeb") },
        {
          url: "https://apps.apple.com/cn/app/%E9%82%BB%E4%BF%A1-%E8%93%9D%E7%89%99%E7%95%85%E8%81%8A/id6472197439",
          label: t("LinChatApp"),
        },
      ],
    },
  ];

  const socialLinks = [
    {
      href: "https://discord.com/channels/1250069935594536960/1250069935594536963",
      ariaLabel: "ZBook on Discord",
      icon: "discord",
    },
    {
      href: "https://github.com/zizdlp/zbook",
      ariaLabel: "ZBook on GitHub",
      icon: "github",
    },
    {
      href: "https://space.bilibili.com/1448262500",
      ariaLabel: "ZBook on Bilibili",
      icon: "bilibili",
    },
    {
      href: "https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg",
      ariaLabel: "ZBook on Youtube",
      icon: "youtube",
    },
  ];

  return (
    <footer className="bg-gradient-to-b from-[#e3d5b1] pt-16  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent text-slate-700 dark:text-slate-500">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div
          className="grid grid-cols-8 gap-4 gap-x-8 px-2 py-8 overflow-x-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
        scrollbar-thumb-slate-200 scrollbar-track-slate-100
        dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
        >
          {sections.map((section, index) => (
            <FooterSection
              key={index}
              title={section.title}
              links={section.links}
            />
          ))}
        </div>

        <div className="flex flex-col items-center border-t-[0.01rem] border-slate-400 dark:border-slate-700 py-10 sm:flex-row-reverse sm:justify-between">
          <div className="flex gap-x-6">
            {socialLinks.map(({ href, ariaLabel, icon }, index) => (
              <Link
                key={index}
                className="group"
                aria-label={ariaLabel}
                href={href}
              >
                <IconItem
                  iconName={icon}
                  className="h-5 w-5 fill-slate-500 group-hover:fill-slate-700"
                />
              </Link>
            ))}
          </div>
          {beian && (
            <Link
              href="https://beian.miit.gov.cn/"
              className="mt-6 text-sm sm:mt-0"
            >
              {beian}
            </Link>
          )}
          <p className="text-sm">{t("Copyright")}</p>
        </div>
      </div>
    </footer>
  );
}
