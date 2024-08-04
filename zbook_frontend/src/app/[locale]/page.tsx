import MotionBounce from "../../components/MotionBounce";
import { FaUsers } from "react-icons/fa";
import { MdVisibility } from "react-icons/md";
import { FaGithub } from "react-icons/fa";
import HomeFooter from "./HomeFooter";
import TabGroup from "./TabGroup";
import { MdChevronRight, MdTableChart } from "react-icons/md";
import { TbMathFunction } from "react-icons/tb";
import { AiOutlineCode } from "react-icons/ai";
import { MdOutlineFeaturedPlayList } from "react-icons/md";

import { Link } from "@/navigation";
import FeatureTabGroup from "./FeatureTabGroup";
import { getTranslations } from "next-intl/server";

import { redirect } from "@/navigation";
import { auth } from "@/auth";
export default async function Home() {
  const t = await getTranslations("HomePage");
  const session = await auth();
  if (session && session.access_token) {
    redirect(`/workspace/${session.username}`); // Navigate to the new post page
  }
  return (
    <div>
      <section>
        <MotionBounce direction="y">
          <div className="relative pt-10 bg-gradient-to-t from-[#abdbed]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent background-animate text-center">
            <div className="relative z-10">
              <div className="px-6 py-16">
                <div className="mx-auto max-w-[22rem] md:max-w-[40rem] font-inter text-[2rem] font-bold leading-tight md:text-6xl md:leading-[1.08] space-y-2 ">
                  <h1>{t("HomeSloganA")}</h1>
                  <h1>{t("HomeSloganB")}</h1>
                </div>
                <p className="mt-4 max-w-[26rem] mx-auto md:max-w-xl text-primary dark:text-primary-light/80 md:leading-loose">
                  {t("HomeSubTitle")}
                </p>
                <div className="px-12 mt-8 flex justify-center items-center flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-10">
                  <Link
                    className="bg-[#65b1e8] dark:bg-transparent dark:border dark:border-[#65b1e8]/[0.08] w-full sm:w-fit bg-opacity-20 text-white rounded-full p-1 flex items-center justify-center"
                    href="/auth/login"
                  >
                    <span className="bg-[#65b1e8] dark:bg-[#65b1e8]/[0.12] hover:bg-[#65b1e8]/80 dark:hover:bg-[#65b1e8]/20 dark:text-primary-light w-full sm:w-fit rounded-full pr-5 pl-8 py-2 flex items-center justify-center space-x-1">
                      <span>{t("SignIn")}</span>
                      <MdChevronRight className="h-5 w-5" />
                    </span>
                  </Link>
                  <Link
                    className="flex items-center space-x-1 hover:text-slate-800 dark:hover:text-gray-100"
                    href={`/workspace/${process.env.DOC_USERNAME}/o/${process.env.DOC_REPONAME}`}
                  >
                    <span>{t("Docs")}</span>
                    <MdChevronRight className="h-5 w-5" />
                  </Link>
                </div>
                <TabGroup />
              </div>
            </div>
          </div>
        </MotionBounce>
      </section>

      <section id="dashboard">
        <MotionBounce direction="y">
          <div className="relative pt-10 bg-gradient-to-t from-[#adeadc]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent background-animate text-center">
            <div className="relative z-10">
              <div className="px-6 py-16">
                <div className="mx-auto max-w-[22rem] md:max-w-[40rem] font-inter text-[2rem] font-bold leading-tight md:text-6xl md:leading-[1.08] space-y-2 ">
                  <h1>{t("ManageUsers")}</h1>
                  <h1>{t("HomeSloganB")}</h1>
                </div>
                <p className="mt-4 max-w-[26rem] mx-auto md:max-w-xl text-primary dark:text-primary-light/80 md:leading-loose">
                  {t("HomeSubTitle")}
                </p>

                <FeatureTabGroup
                  categories={[t("DashBoardDemo"), t("DashBoardDemoVideo")]}
                  image_urls={["/admin_light.png", "/admin_dark.png"]}
                  video_urls={[
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624008170989&bvid=BV12PV5eCE8F&cid=500001583709871&p=1&autoplay=false",
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624008233402&bvid=BV1TPV5eyE9B&cid=500001583708862&p=1&autoplay=false",
                  ]}
                />
              </div>
            </div>
          </div>
        </MotionBounce>
      </section>

      <section id="create_repo">
        <MotionBounce direction="y">
          <div className="relative pt-10  bg-gradient-to-t from-[#adeab8]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent background-animate text-center">
            <div className="relative z-10">
              <div className="px-6 py-16">
                <div className="mx-auto max-w-[22rem] md:max-w-[40rem] font-inter text-[2rem] font-bold leading-tight md:text-6xl md:leading-[1.08] space-y-2">
                  <h1>{t("NewRepo")}</h1>
                  <h1>{t("HomeSloganB")}</h1>
                </div>
                <p className="mt-4 max-w-[26rem] mx-auto md:max-w-xl text-primary dark:text-primary-light/80 md:leading-loose">
                  {t("HomeSubTitle")}
                </p>

                <FeatureTabGroup
                  categories={[t("DashBoardDemo"), t("DashBoardDemoVideo")]}
                  image_urls={[
                    "/create_repo_light.png",
                    "/create_repo_dark.png",
                  ]}
                  video_urls={[
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624025013827&bvid=BV1TAV5eLEmr&cid=500001583711218&p=1&autoplay=false",
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624024948260&bvid=BV1mAV5e5Ena&cid=500001583710383&p=1&autoplay=false",
                  ]}
                />
              </div>
            </div>
          </div>
        </MotionBounce>
      </section>

      <section id="features_section">
        <MotionBounce direction="y">
          <div className="relative pt-10  bg-gradient-to-t from-[#d5eaad]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent background-animate text-center">
            <div className="relative z-10">
              <div className="px-6 py-16">
                <div className="mx-auto max-w-[22rem] md:max-w-[40rem] font-inter text-[2rem] font-bold leading-tight md:text-6xl md:leading-[1.08] space-y-2">
                  <h1>{t("MarkdownSuperset")}</h1>
                  <h1>{t("HomeSloganB")}</h1>
                </div>
                <p className="mt-4 max-w-[26rem] mx-auto md:max-w-xl text-primary dark:text-primary-light/80 md:leading-loose">
                  {t("HomeSubTitle")}
                </p>

                <FeatureTabGroup
                  categories={[t("DashBoardDemo"), t("DashBoardDemoVideo")]}
                  image_urls={["/feature_light.png", "/feature_dark.png"]}
                  video_urls={[
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624025013879&bvid=BV1TAV5eLEUw&cid=500001583712129&p=1&autoplay=false",
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624025010780&bvid=BV1MAV5eLEkH&cid=500001583711739&p=1&autoplay=false",
                  ]}
                />
              </div>
            </div>
          </div>
        </MotionBounce>
      </section>

      <section id="multi_user_section">
        <MotionBounce direction="y">
          <div className="relative pt-10 bg-gradient-to-t from-[#96c4a5]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent background-animate text-center">
            <div className="relative z-10">
              <div className="px-6 py-16">
                <div className="mx-auto max-w-[22rem] md:max-w-[40rem] font-inter text-[2rem] font-bold leading-tight md:text-6xl md:leading-[1.08] space-y-2 ">
                  <h1>{t("MultiUserSection")}</h1>
                  <h1>{t("HomeSloganB")}</h1>
                </div>
                <p className="mt-4 max-w-[26rem] mx-auto md:max-w-xl text-primary dark:text-primary-light/80 md:leading-loose">
                  {t("HomeSubTitle")}
                </p>
                <FeatureTabGroup
                  categories={[t("DashBoardDemo"), t("DashBoardDemoVideo")]}
                  image_urls={["/session_light.png", "/session_dark.png"]}
                  video_urls={[
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624025141935&bvid=BV1UAV5eLEz9&cid=500001583712733&p=1&autoplay=false",
                    "https://player.bilibili.com/player.html?isOutside=true&aid=112624025013907&bvid=BV1TAV5eLE2E&cid=500001583712437&p=1&autoplay=false",
                  ]}
                />
              </div>
            </div>
          </div>
        </MotionBounce>
      </section>

      <section className=" py-20 sm:py-32 bg-gradient-to-t from-[#e3d5b1]  to-[#ffffff] dark:from-transparent dark:via-transparent dark:to-transparent">
        <MotionBounce direction="y">
          <div className="my-12 px-6 md:px-0 mx-auto max-w-5xl text-center">
            <div className="w-full text-center">
              <div className="flex items-center justify-center space-x-1.5 text-sm">
                <MdOutlineFeaturedPlayList />
                <span>{t("FeatureSection")}</span>
              </div>
              <div className="mt-6 font-bold font-inter text-[1.7rem] leading-[1.3] md:text-5xl items-center justify-center md:leading-[1.15]  flex flex-col">
                <h1 className="max-w-2xl block">{t("SimpleDocumentation")}</h1>
                <h1 className="max-w-2xl block">{t("VisuallyAppealing")}</h1>
              </div>
              <div className="mt-3.5 flex justify-center">
                <h3 className="max-w-sm">{t("DocumentationForTeam")}</h3>
              </div>
            </div>
            <div className="px-12 mt-8 flex justify-center items-center flex-col sm:flex-row space-y-4 sm:space-y-0 sm:space-x-10">
              <Link
                className="bg-[#65b1e8] dark:bg-transparent dark:border dark:border-[#65b1e8]/[0.08] w-full sm:w-fit bg-opacity-20 text-white rounded-full p-1 flex items-center justify-center"
                href="/cases"
              >
                <span className="bg-[#65b1e8] dark:bg-[#65b1e8]/[0.12] hover:bg-[#65b1e8]/80 dark:hover:bg-[#65b1e8]/20 dark:text-primary-light w-full sm:w-fit rounded-full pr-5 pl-8 py-2 flex items-center justify-center space-x-1">
                  <span>{t("Cases")}</span>
                  <MdChevronRight className="h-5 w-5" />
                </span>
              </Link>
              <Link
                className=" flex items-center space-x-1 hover:text-slate-800 dark:hover:text-gray-100"
                href={`/workspace/${process.env.DOC_USERNAME}/o/${process.env.DOC_REPONAME}`}
              >
                <span>{t("Docs")}</span>
                <MdChevronRight className="h-5 w-5" />
              </Link>
            </div>
          </div>

          <div className="mx-auto max-w-6xl grid md:grid-cols-3 text-left bg-gray-900/5 dark:bg-white/[0.07] gap-px border-x border-[#EBEBEF] dark:border-[#181A21]">
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <TbMathFunction className="h-8 w-8" />
              <h1 className="mt-3">{t("Math")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("MathDetail")}</h2>
            </div>
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <AiOutlineCode className="h-8 w-8" />
              <h1 className="mt-3">{t("Code")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("CodeDetail")}</h2>
            </div>
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <MdTableChart className="h-8 w-8" />
              <h1 className="mt-3">{t("Figure")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("FigureDetail")}</h2>
            </div>
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <FaUsers className="h-8 w-8" />
              <h1 className="mt-3">{t("MultiUserSection")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("ZBookSupports")}</h2>
            </div>
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <MdVisibility className="h-8 w-8" />
              <h1 className="mt-3">{t("MultiLevelPermissions")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("ZBookPermissions")}</h2>
            </div>
            <div className="px-8 py-7 bg-white dark:bg-slate-900/50">
              <FaGithub className="h-8 w-8" />
              <h1 className="mt-3">{t("SelfHost")}</h1>
              <h2 className="mt-2 leading-relaxed">{t("ZBookIsOpenSource")}</h2>
            </div>
          </div>
        </MotionBounce>
      </section>
      <HomeFooter />
    </div>
  );
}
