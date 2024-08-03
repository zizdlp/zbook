/* eslint-disable react/jsx-no-literals */
import FrpcSideBar from "./FrpcSideBar";

export default function frpc() {
  return (
    <div className="relative antialiased text-[#757077] dark:text-[#A39FA6]">
      <span className="fixed inset-0 bg-background-light dark:bg-background-dark"></span>
      <span className="fixed inset-0"></span>

      <div className="px-4 mx-auto max-w-[92rem] lg:px-8">
        <FrpcSideBar />
        <div className="max-w-6xl" id="content-container">
          <div className="flex flex-row items-stretch gap-12 pt-[3.5rem] lg:pt-[3.2rem]">
            <div
              className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem]"
              id="content-area"
            >
              <header id="header" className="relative">
                <div className="mt-0.5 space-y-2.5">
                  <div className="eyebrow h-5 text-primary dark:text-primary-light text-sm font-semibold">
                    Welcome
                  </div>
                  <div className="flex items-center">
                    <h1 className="inline-block text-2xl sm:text-3xl font-extrabold text-gray-900 tracking-tight dark:text-gray-200">
                      fRPC Documentation
                    </h1>
                  </div>
                </div>
              </header>
              <div className="flex flex-col gap-8">
                <div className="flex flex-col gap-6 xl:hidden"></div>
              </div>
              <div className="relative mt-8 prose prose-gray dark:prose-invert max-w-7xl">
                <p>
                  <button data-state="closed">
                    <span className="underline decoration-dotted decoration-2 underline-offset-4 decoration-gray-400 dark:decoration-gray-500">
                      fRPC
                    </span>
                  </button>{" "}
                  is a <strong>proto3-compatible</strong> RPC Framework that’s
                  designed from the ground up to be{" "}
                  <strong>
                    lightweight, extensible, and extremely performant
                  </strong>
                  . On average fRPC outperforms other RPC frameworks{" "}
                  <a href="/performance/grpc-benchmarks">
                    by 2-4x in an apples-to-apples comparison
                  </a>
                  , and is easily able to handle more than 2 million RPCs/second
                  on a single server.
                </p>
                <span aria-owns="rmiz-modal-3deaf47cc47c" data-rmiz="">
                  <span data-rmiz-content="found">
                    <img
                      className="w-full"
                      src="https://mintlify.s3-us-west-1.amazonaws.com/frpc/images/intro.svg"
                      alt="Welcome to fRPC"
                    />
                  </span>
                  <span data-rmiz-ghost="">
                    <button
                      aria-label="Expand image: Welcome to fRPC"
                      data-rmiz-btn-zoom=""
                      type="button"
                    >
                      <svg
                        aria-hidden="true"
                        data-rmiz-btn-zoom-icon="true"
                        fill="currentColor"
                        focusable="false"
                        viewBox="0 0 16 16"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path d="M 9 1 L 9 2 L 12.292969 2 L 2 12.292969 L 2 9 L 1 9 L 1 14 L 6 14 L 6 13 L 2.707031 13 L 13 2.707031 L 13 6 L 14 6 L 14 1 Z"></path>
                      </svg>
                    </button>
                  </span>
                </span>
                <div className="not-prose grid gap-x-4 sm:grid-cols-2">
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/overview"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Getting Started
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Quickly get up and running with fRPC by following
                            our getting started guide.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/concepts"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Concepts
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Take a look at some unique fRPC concepts and how it
                            differs from other frameworks.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/getting-started/concepts"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Roadmap
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Check out our planned technical roadmap to see how
                            we’ll be improving fRPC in the future.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    href="/reference/overview"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Technical Docs
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            Take a look at our technical docs to dig into the
                            details of fRPC and how you can use it.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                </div>
                <div className="not-prose grid gap-x-4 sm:grid-cols-2">
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    target="_blank"
                    rel="noreferrer"
                    href="https://github.com/loopholelabs/frpc-go"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Contributing
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            The footer of each page contains an “Edit on Github”
                            link. Please feel free to contribute by making a
                            pull request!
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                  <a
                    className="card block not-prose font-normal group relative my-2 ring-2 ring-transparent rounded-xl bg-white dark:bg-background-dark border border-gray-950/10 dark:border-white/10 overflow-hidden w-full cursor-pointer hover:!border-primary dark:hover:!border-primary-light"
                    target="_blank"
                    rel="noreferrer"
                    href="https://loopholelabs.io/discord"
                  >
                    <div className="px-6 py-5">
                      <div className="h-6 w-6 fill-gray-800 dark:fill-gray-100 text-gray-800 dark:text-gray-100">
                        <svg className="h-6 w-6 bg-primary dark:bg-primary-light"></svg>
                      </div>
                      <div>
                        <h2 className="font-semibold text-base text-gray-800 dark:text-white mt-4">
                          Join Our Discord
                        </h2>
                        <div className="mt-1 font-normal text-sm leading-6 text-gray-600 dark:text-gray-400">
                          <p>
                            The <strong>#Frisbee</strong> Channel in our Discord
                            server is a great place to get help with all things
                            Frisbee and fRPC.
                          </p>
                        </div>
                      </div>
                    </div>
                  </a>
                </div>
              </div>
              <div className="leading-6 mt-14">
                <div className="mb-12 px-0.5 flex items-center text-sm font-semibold text-gray-700 dark:text-gray-200">
                  <a
                    className="flex items-center ml-auto space-x-3 group"
                    href="/getting-started/overview"
                  >
                    <span className="group-hover:text-gray-900 dark:group-hover:text-white">
                      Overview
                    </span>
                    <svg
                      viewBox="0 0 3 6"
                      className="rotate-180 h-1.5 stroke-gray-400 overflow-visible group-hover:stroke-gray-600 dark:group-hover:stroke-gray-300"
                    >
                      <path
                        d="M3 0L0 3L3 6"
                        fill="none"
                        stroke-width="2"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      ></path>
                    </svg>
                  </a>
                </div>
                <footer className="justify-between pt-10 border-t border-gray-100 sm:flex dark:border-gray-800/50 pb-28">
                  <div className="flex mb-6 space-x-6 sm:mb-0">
                    <a href="https://loopholelabs.io/discord" target="_blank">
                      <span className="sr-only">discord</span>
                      <svg className="w-5 h-5 bg-gray-400 dark:bg-gray-500 hover:bg-gray-500 dark:hover:bg-gray-400"></svg>
                    </a>
                    <a
                      href="https://github.com/loopholelabs/frpc-go"
                      target="_blank"
                    >
                      <span className="sr-only">github</span>
                      <svg className="w-5 h-5 bg-gray-400 dark:bg-gray-500 hover:bg-gray-500 dark:hover:bg-gray-400"></svg>
                    </a>
                  </div>
                  <div className="sm:flex">
                    <a
                      href="https://mintlify.com?utm_campaign=poweredBy&amp;utm_medium=docs&amp;utm_source=frpc.io"
                      target="_blank"
                      rel="noreferrer"
                      className="text-sm text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-300"
                    >
                      Powered by Mintlify
                    </a>
                  </div>
                </footer>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
