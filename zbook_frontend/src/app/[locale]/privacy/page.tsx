/* eslint-disable react/jsx-no-literals */
export default function Privacy() {
  return (
    <div className="container mx-auto px-4 py-8 max-w-3xl dark:text-gray-400 text-gray-800">
      <h1 className="text-4xl font-bold mb-6">Privacy Policy</h1>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">1. Introduction</h2>
        <p className="mb-4">
          We are committed to protecting your privacy. This Privacy Policy
          explains how we collect, use, and disclose information when you use
          <strong className="px-2">ZBook</strong>.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">
          2. Information We Collect
        </h2>
        <p className="mb-4">
          <strong>Personal Information:</strong> We do not collect any personal
          information unless you voluntarily provide it to us, such as by
          contacting us via email or signing up for our newsletter.
        </p>
        <p className="mb-4">
          <strong>Usage Data:</strong> We may collect information on how the
          Software is accessed a nd used. This usage data may include
          information such as your computers Internet Protocol IP address,
          browser type, browser version, the pages of our Software that you
          visit, the time and date of your visit, the time spent on those pages,
          and other diagnostic data.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">3. Use of Data</h2>
        <p className="mb-4">We use the collected data for various purposes:</p>
        <ul className="list-disc list-inside mb-4">
          <li>To provide and maintain the Software</li>
          <li>To notify you about changes to our Software</li>
          <li>To allow you to participate in interactive features</li>
        </ul>
      </section>
    </div>
  );
}
