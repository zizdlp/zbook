/* eslint-disable react/jsx-no-literals */
import { Link } from "@/navigation";
export default function Terms() {
  return (
    <div className="container mx-auto px-4 py-8 max-w-3xl dark:text-gray-400 text-gray-800">
      <h1 className="text-4xl font-bold mb-6">Terms of Service</h1>
      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">1. Acceptance of Terms</h2>
        <p className="mb-4">
          By accessing and using <strong>ZBook</strong>, you agree to comply
          with and be bound by these Terms of Service (Terms). If you do not
          agree to these Terms, please do not use the Software.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">2. License</h2>
        <p className="mb-4">
          <strong>ZBook</strong> is licensed under the{" "}
          <strong className="text-sky-500 hover:text-sky-700">
            <Link href="https://github.com/zizdlp/mdwiki?tab=GPL-3.0-1-ov-file#readme">
              GPL-3.0 license
            </Link>
          </strong>
          . This means you are free to use, copy of the Software, subject to the
          conditions outlined in the license.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">3. Use of Software</h2>
        <p className="mb-4">
          You agree to use the Software in compliance with all applicable laws
          and regulations. You are responsible for ensuring that your use of the
          Software does not violate any rights of third parties.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">
          4. Modifications to the Software
        </h2>
        <p className="mb-4">
          We reserve the right to modify, suspend, or discontinue the Software
          at any time without notice. We are not liable to you or any third
          party for any modifications, suspensions, or discontinuations of the
          Software.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">
          5. Disclaimer of Warranties
        </h2>
        <p className="mb-4">
          The Software is provided as is, without warranty of any kind, express
          or implied, including but not limited to the warranties of
          merchantability, fitness for a particular purpose, and
          noninfringement. In no event shall the authors or copyright holders be
          liable for any claim, damages, or other liability, whether in an
          action of contract, tort, or otherwise, arising from, out of, or in
          connection with the Software or the use or other dealings in the
          Software.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">
          6. Limitation of Liability
        </h2>
        <p className="mb-4">
          In no event shall we be liable for any indirect, incidental, special,
          consequential, or punitive damages, or any loss of profits or
          revenues, whether incurred directly or indirectly, or any loss of
          data, use, goodwill, or other intangible losses, resulting from (i)
          your use or inability to use the Software; (ii) any unauthorized
          access to or use of our servers and/or any personal information stored
          therein; (iii) any interruption or cessation of transmission to or
          from the Software; (iv) any bugs, viruses, trojan horses, or the like
          that may be transmitted to or through the Software by any third party;
          (v) any errors or omissions in any content or for any loss or damage
          incurred as a result of the use of any content posted, emailed,
          transmitted, or otherwise made available through the Software; and/or
          (vi) the defamatory, offensive, or illegal conduct of any third party.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">7. Governing Law</h2>
        <p className="mb-4">
          These Terms shall be governed and construed in accordance with the
          laws of [Your Country/State], without regard to its conflict of law
          provisions.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">8. Changes to the Terms</h2>
        <p className="mb-4">
          We reserve the right, at our sole discretion, to modify or replace
          these Terms at any time. If a revision is material, we will provide at
          least 30 days notice prior to any new terms taking effect. What
          constitutes a material change will be determined at our sole
          discretion.
        </p>
      </section>

      <section className="mb-8">
        <h2 className="text-2xl font-semibold mb-4">9. Contact Us</h2>
        <p className="mb-4">
          If you have any questions about these Terms, please contact us at
          zizdlp@gmail.com.
        </p>
      </section>
    </div>
  );
}
