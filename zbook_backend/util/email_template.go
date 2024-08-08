package util

const EmailTemplate = `<!DOCTYPE html>
<html lang="en">
  <body
    style="
      font-family: Arial, sans-serif;
      background-color: #f7f7f7;
      margin: 0;
      padding: 0;
    "
  >
    <div
      style="
        max-width: 768px;
        margin: 40px auto;
        background-color: white;
        padding: 24px;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      "
    >
      <div
        style="
          text-align: center;
          border-bottom: 1px solid #e5e7eb;
          padding-bottom: 16px;
          margin-bottom: 16px;
        "
      >
        <h1
          style="font-size: 1.5rem; font-weight: bold; color: black; margin: 0"
        >
          %s
        </h1>
      </div>
      <div style="margin-top: 16px">
        <p style="font-size: 1.125rem; margin: 0 0 16px">Hello, %s</p>
        <p style="font-size: 1.125rem; margin: 0 0 16px">
          %s
        </p>
        <a
          href="%s"
          target="_blank"
          style="
            display: inline-block;
            background-color: #2563eb;
            color: white;
            padding: 8px 16px;
            border-radius: 4px;
            font-weight: bold;
            text-decoration: none;
            transition: background-color 0.3s;
          "
          onmouseover="this.style.backgroundColor='#1d4ed8';"
          onmouseout="this.style.backgroundColor='#2563eb';"
        >
         %s
        </a>
        <p style="font-size: 1.125rem; margin: 16px 0">
         %s
        </p>
        <p style="font-size: 1.125rem; margin: 16px 0">
          Thanks,<br />The ZBook Team
        </p>
      </div>
      <div
        style="
          text-align: center;
          font-size: 14px;
          color: #888888;
          margin-top: 20px;
          padding: 20px;
          border-top: 1px solid #eaeaea;
        "
      >
        <table align="center" style="margin: 0 auto">
          <tr>
            <td style="padding-right: 10px; vertical-align: middle">
              <img
                src="data:image/png;base64,%s"
                alt="LOGO"
                style="
                  width: 40px;
                  height: 40px;
                  border-radius: 8px;
                  margin-bottom: 10px;
                "
              />
            </td>
            <td style="vertical-align: middle">
              <div style="font-size: 16px; font-weight: 600">ZBook</div>
            </td>
          </tr>
        </table>
        <div style="margin: 20px 0">
          <a
            href="https://github.com/zizdlp/zbook"
            style="margin: 0 10px; text-decoration: none; color: #3b82f6"
            >GitHub</a
          >
          <a
            href="https://space.bilibili.com/1448262500"
            style="margin: 0 10px; text-decoration: none; color: #3b82f6"
            >Bilibili</a
          >
          <a
            href="https://www.youtube.com/channel/UC9D6VAJRoG7bD38dz8F9CSg"
            style="margin: 0 10px; text-decoration: none; color: #3b82f6"
            >YouTube</a
          >
          <a
            href="https://discord.com/channels/1250069935594536960/1250069935594536963"
            style="margin: 0 10px; text-decoration: none; color: #3b82f6"
            >Discord</a
          >
        </div>
        <p style="font-size: 14px; color: #888888">
          © 2024 zizdlp.com. All rights reserved.
        </p>
      </div>
    </div>
  </body>
</html>`
