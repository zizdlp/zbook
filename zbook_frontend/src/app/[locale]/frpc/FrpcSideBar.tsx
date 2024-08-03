/* eslint-disable react/jsx-no-literals */
import SideBarSearchButton from "@/components/sidebars/SideBarSearchButton";
import { SearchType } from "@/utils/const_value";
import FrpcAnchor from "./FrpcAnchor";

export default function FrpcSideBar() {
  const layout = [
    {
      title: "Welcome",
      relative_path: "",
      isdir: true,
      sublayouts: [
        {
          title: "introduction",
          relative_path: "introduction",
          isdir: false,
          sublayouts: null,
        },
      ],
    },
    {
      title: "Getting Started",
      relative_path: "getting-started",
      isdir: true,
      sublayouts: [
        {
          title: "overview",
          relative_path: "getting-started/overview",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "quick-start",
          relative_path: "getting-started/quick-start",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "concepts",
          relative_path: "getting-started/concepts",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "architecture",
          relative_path: "getting-started/architecture",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "architecture",
          relative_path: "getting-started/roadmap",
          isdir: false,
          sublayouts: null,
        },
      ],
    },
    {
      title: "Performance",
      relative_path: "performance",
      isdir: true,
      sublayouts: [
        {
          title: "optimizations",
          relative_path: "performance/optimizations",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "grpc-benchmarks",
          relative_path: "performance/grpc-benchmarks",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "twirp-benchmarks",
          relative_path: "performance/twirp-benchmarks",
          isdir: false,
          sublayouts: null,
        },
      ],
    },
    {
      title: "Reference",
      relative_path: "reference",
      isdir: true,
      sublayouts: [
        {
          title: "overview",
          relative_path: "reference/overview",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "client-methods",
          relative_path: "reference/client-methods",
          isdir: false,
          sublayouts: null,
        },
        {
          title: "server-methods",
          relative_path: "reference/server-methods",
          isdir: false,
          sublayouts: null,
        },
      ],
    },
  ];
  return (
    <div
      className="z-20 hidden lg:block fixed bottom-0 right-auto w-[18rem] top-[4rem]"
      id="sidebar"
    >
      <div
        className="absolute inset-0 z-10 overflow-auto pr-8 pb-10"
        id="sidebar-content"
      >
        <div className="relative lg:text-sm lg:leading-6">
          <div className="sticky top-0 pointer-events-none">
            <div className="h-10 bg-white dark:bg-gray-900"></div>
            <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
              <SideBarSearchButton
                username={"a"}
                repo_name={"b"}
                searchType={SearchType.REPO_DOCUMENT}
              />
            </div>
          </div>
          <div className="sticky top-0 h-8"></div>
          <ul id="navigation-items">
            <FrpcAnchor />
            {layout.map((item, index) => (
              <div key={index} className="mt-12 lg:mt-8">
                <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                  {item.title}
                </h5>
                {item.sublayouts.map((subitem, index) => (
                  <li key={index}>
                    <a
                      className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg bg-primary/10 text-primary font-semibold dark:text-primary-light dark:bg-primary-light/10"
                      href="/introduction"
                    >
                      <div className="flex-1 flex items-center space-x-2.5">
                        <div>{subitem.title}</div>
                      </div>
                    </a>
                  </li>
                ))}
              </div>
            ))}

            <div className="mt-12 lg:mt-8">
              <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                Welcome
              </h5>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg bg-primary/10 text-primary font-semibold dark:text-primary-light dark:bg-primary-light/10"
                  href="/introduction"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Introduction</div>
                  </div>
                </a>
              </li>
            </div>
            <div className="mt-12 lg:mt-8">
              <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                Getting Started
              </h5>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/getting-started/overview"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Overview</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/getting-started/quick-start"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Quick Start</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/getting-started/concepts"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Concepts</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/getting-started/architecture"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Architecture</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/getting-started/roadmap"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Roadmap</div>
                  </div>
                </a>
              </li>
            </div>
            <div className="mt-12 lg:mt-8">
              <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                Performance
              </h5>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/performance/optimizations"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Optimizations</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/performance/grpc-benchmarks"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>gRPC Benchmarks</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/performance/twirp-benchmarks"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Twirp Benchmarks</div>
                  </div>
                </a>
              </li>
            </div>
            <div className="mt-12 lg:mt-8">
              <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                Reference
              </h5>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/reference/overview"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Overview</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/reference/client-methods"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Client Methods</div>
                  </div>
                </a>
              </li>
              <li>
                <a
                  className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                  href="/reference/server-methods"
                >
                  <div className="flex-1 flex items-center space-x-2.5">
                    <div>Server Methods</div>
                  </div>
                </a>
              </li>
            </div>
          </ul>
        </div>
      </div>
    </div>
  );
}
