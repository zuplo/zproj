import { Head, Link } from "zudoku/components";
import { useEffect, useRef, useState } from "react";

const steps = [
  {
    label: "Setup",
    lines: [
      { type: "cmd", text: "hk init" },
      { type: "out", text: "Created hike.yaml" },
      { type: "gap" },
      { type: "comment", text: "# Edit hike.yaml to add your repos..." },
      { type: "gap" },
      { type: "cmd", text: "hk sync" },
      { type: "out", text: 'Syncing group "platform" (3 repos)...' },
      { type: "out", text: "  my-app:     cloned", delay: 300 },
      { type: "out", text: "  shared-lib: cloned", delay: 150 },
      { type: "out", text: "  api:        cloned", delay: 200 },
      { type: "out", text: "Sync complete." },
    ],
  },
  {
    label: "Create",
    lines: [
      { type: "cmd", text: "hk platform" },
      { type: "out", text: "Generated project name: bold-cedar" },
      {
        type: "out",
        text: 'Creating project "platform-bold-cedar" in group "platform"...',
      },
      { type: "out", text: "Running onCreate hooks..." },
      { type: "out", text: "  my-app:     hook done", delay: 400 },
      { type: "out", text: "  shared-lib: hook done", delay: 200 },
      { type: "out", text: "  api:        hook done", delay: 300 },
      {
        type: "success",
        text: 'Project "platform-bold-cedar" created successfully.',
      },
      { type: "gap" },
      { type: "cmd", text: "ls platform-bold-cedar/" },
      {
        type: "out",
        text: "my-app/  shared-lib/  api/  platform-bold-cedar.code-workspace",
      },
    ],
  },
  {
    label: "Work",
    lines: [
      { type: "cmd", text: "cd platform-bold-cedar" },
      { type: "gap" },
      { type: "cmd", text: "hk status" },
      { type: "out", text: "Project: platform-bold-cedar" },
      { type: "out", text: "" },
      { type: "out", text: "  my-app       branch: platform-bold-cedar" },
      {
        type: "out",
        text: "  shared-lib   branch: platform-bold-cedar  [dirty]",
      },
      { type: "out", text: "  api          branch: platform-bold-cedar" },
      { type: "gap" },
      { type: "cmd", text: "hk pull" },
      { type: "out", text: "  my-app:     pulled" },
      { type: "out", text: "  shared-lib: pulled" },
      { type: "out", text: "  api:        pulled" },
    ],
  },
  {
    label: "Clean up",
    lines: [
      { type: "cmd", text: "hk delete platform-bold-cedar" },
      { type: "out", text: 'Deleting project "platform-bold-cedar"...' },
      { type: "out", text: "  my-app:     removed", delay: 200 },
      { type: "out", text: "  shared-lib: removed", delay: 150 },
      { type: "out", text: "  api:        removed", delay: 180 },
      { type: "success", text: 'Project "platform-bold-cedar" deleted.' },
      { type: "gap" },
      { type: "cmd", text: "hk list" },
      { type: "out", text: "No projects found." },
    ],
  },
];

function wait(ms: number) {
  return new Promise((r) => setTimeout(r, ms));
}

function TerminalDemo() {
  const [current, setCurrent] = useState(0);
  const [lines, setLines] = useState<
    { type: string; text: string; typed?: string }[]
  >([]);
  const bodyRef = useRef<HTMLDivElement>(null);
  const playingRef = useRef(false);

  useEffect(() => {
    playStep(current);
  }, [current]);

  async function playStep(idx: number) {
    if (playingRef.current) return;
    playingRef.current = true;
    setLines([]);
    const step = steps[idx];

    for (const line of step.lines) {
      if (line.type === "gap") {
        setLines((prev) => [...prev, { type: "gap", text: "" }]);
        await wait(100);
        continue;
      }
      if (line.type === "cmd") {
        // Type character by character
        const text = line.text;
        for (let i = 0; i <= text.length; i++) {
          const partial = text.slice(0, i);
          setLines((prev) => {
            const next = [...prev];
            const last = next[next.length - 1];
            if (last?.type === "cmd-typing") {
              last.typed = partial;
              return [...next];
            }
            return [...next, { type: "cmd-typing", text, typed: partial }];
          });
          await wait(35 + Math.random() * 20);
        }
        // Finalize
        setLines((prev) => {
          const next = [...prev];
          next[next.length - 1] = { type: "cmd", text };
          return next;
        });
        await wait(300);
      } else {
        await wait(line.delay || 80);
        setLines((prev) => [...prev, { type: line.type, text: line.text }]);
      }
      bodyRef.current?.scrollTo(0, bodyRef.current.scrollHeight);
    }

    playingRef.current = false;
    await wait(3000);
    if (idx === current) {
      setCurrent((idx + 1) % steps.length);
    }
  }

  return (
    <div>
      <div className="rounded-xl border border-[#172028] bg-[#0a1016] overflow-hidden max-w-3xl mx-auto">
        <div className="flex items-center gap-2 px-4 py-3 border-b border-[#172028]">
          <div className="w-3 h-3 rounded-full bg-[#ff5f57]" />
          <div className="w-3 h-3 rounded-full bg-[#febc2e]" />
          <div className="w-3 h-3 rounded-full bg-[#28c840]" />
          <div className="flex-1 text-center text-xs text-[#7e8f9e] font-mono">
            ~/projects
          </div>
        </div>
        <div
          ref={bodyRef}
          className="p-5 min-h-[320px] max-h-[380px] overflow-y-auto font-mono text-sm leading-[1.8]"
        >
          {lines.map((line, i) => {
            if (line.type === "gap") return <div key={i} className="h-2" />;
            if (line.type === "cmd-typing")
              return (
                <div key={i}>
                  <span className="text-green-500">$ </span>
                  {line.typed}
                  <span className="text-green-400 animate-pulse">|</span>
                </div>
              );
            if (line.type === "cmd")
              return (
                <div key={i}>
                  <span className="text-green-500">$ </span>
                  {line.text}
                </div>
              );
            if (line.type === "comment")
              return (
                <div key={i} className="text-[#3e5260]">
                  {line.text}
                </div>
              );
            if (line.type === "success")
              return (
                <div key={i} className="text-green-400 font-medium">
                  {line.text}
                </div>
              );
            return (
              <div key={i} className="text-[#7e8f9e]">
                {line.text}
              </div>
            );
          })}
        </div>
      </div>
      <div className="flex gap-2 justify-center mt-5">
        {steps.map((step, i) => (
          <button
            key={i}
            onClick={() => {
              playingRef.current = false;
              setCurrent(i);
            }}
            className={`px-4 py-1.5 rounded-full text-sm font-semibold border transition-all cursor-pointer ${
              current === i
                ? "bg-green-500/10 border-green-500/25 text-green-400"
                : "bg-white/[0.03] border-[#172028] text-[#7e8f9e] hover:text-white hover:border-[#1e2d38]"
            }`}
          >
            {step.label}
          </button>
        ))}
      </div>
    </div>
  );
}

export const HomePage = () => {
  return (
    <div className="bg-[#050a0e] text-[#e4e9ee] min-h-screen">
      <Head>
        <title>hike — Multi-repo workspaces for AI workflows</title>
      </Head>

      {/* Nav */}
      <nav className="fixed top-0 left-0 right-0 z-50 py-3.5 bg-[#050a0e]/85 backdrop-blur-xl border-b border-[#172028]">
        <div className="max-w-5xl mx-auto px-6 flex items-center justify-between">
          <a href="/hike/" className="flex items-center gap-2.5 font-extrabold text-[#e4e9ee] no-underline hover:no-underline">
            <svg className="w-10 h-10" viewBox="0 0 128 128" fill="none"><path d="M12 104 L52 32 L64 52 L76 28 L116 104 Z" fill="#15803d"/><path d="M64 104 L68 76 L60 60 L68 44 L76 28" stroke="#fbbf24" strokeWidth="4" strokeLinecap="round" fill="none"/></svg>
            <span className="text-xl">hike</span>
          </a>
          <div className="flex items-center gap-6">
            <Link to="/overview" className="text-sm font-semibold text-[#e4e9ee] no-underline hover:text-green-400 hover:no-underline">Documentation</Link>
            <a href="https://github.com/zuplo/hike" className="inline-flex items-center gap-1.5 bg-white/[0.06] border border-[#1e2d38] rounded-lg px-3.5 py-1.5 text-sm font-medium text-[#e4e9ee] no-underline hover:border-green-500/30 hover:text-green-400 hover:no-underline transition-all">
              <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor"><path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"/></svg>
              GitHub
            </a>
          </div>
        </div>
      </nav>

      {/* Hero */}
      <section className="text-center pt-32 pb-20 px-6 max-w-4xl mx-auto">
        <div className="inline-flex items-center gap-2 bg-green-500/[0.07] border border-green-500/20 rounded-full px-4 py-1 text-xs font-semibold text-green-400 mb-9">
          <svg
            width="14"
            height="14"
            viewBox="0 0 16 16"
            fill="currentColor"
          >
            <path d="M8 0l2.5 5.3L16 6.2l-4 3.9.9 5.9L8 13.4 3.1 16l.9-5.9-4-3.9 5.5-.9z" />
          </svg>
          Built for AI-powered development
        </div>
        <h1 className="text-5xl md:text-7xl font-black tracking-tighter leading-[1.08] mb-6 bg-gradient-to-br from-white to-green-400 bg-clip-text text-transparent">
          Multi-repo workspaces
          <br />
          for AI workflows.
        </h1>
        <p className="text-lg text-[#7e8f9e] max-w-xl mx-auto mb-12 leading-relaxed">
          Give every AI coding agent its own isolated workspace across all your
          repos. Git worktrees mean zero conflicts, instant setup, and clean
          teardown.
        </p>

        <div className="text-[0.7rem] font-bold uppercase tracking-widest text-[#7e8f9e] mb-2">
          Install
        </div>
        <div className="inline-flex items-center bg-[#0a1016] border border-[#172028] rounded-xl px-6 py-4 font-mono text-sm overflow-x-auto max-w-full">
          <span className="text-green-500 mr-2">$</span>
          <span>
            curl{" "}
            <span className="text-yellow-400">-fsSL</span>{" "}
            https://raw.githubusercontent.com/zuplo/hike/main/install.sh | sh
          </span>
        </div>

        <div className="flex gap-3 justify-center mt-8 flex-wrap">
          <Link
            to="/overview"
            className="inline-flex items-center gap-2 px-7 py-3.5 rounded-xl font-semibold bg-green-600 text-white shadow-[0_0_30px_rgba(34,197,94,0.2)] hover:bg-green-500 transition-all hover:no-underline hover:-translate-y-0.5"
          >
            Get Started
          </Link>
          <a
            href="https://github.com/zuplo/hike"
            className="inline-flex items-center gap-2 px-7 py-3.5 rounded-xl font-semibold bg-white/[0.04] border border-[#172028] text-[#e4e9ee] hover:border-[#1e2d38] transition-all hover:no-underline hover:-translate-y-0.5"
          >
            GitHub
          </a>
        </div>
      </section>

      <div className="h-px bg-[#172028]" />

      {/* Demo */}
      <section className="py-24 px-6 max-w-4xl mx-auto">
        <div className="text-center mb-12">
          <div className="text-[0.72rem] font-bold uppercase tracking-widest text-green-400 mb-3">
            How it works
          </div>
          <h2 className="text-3xl md:text-4xl font-extrabold tracking-tight mb-4">
            From zero to isolated workspace in seconds.
          </h2>
          <p className="text-[#7e8f9e] text-lg max-w-lg mx-auto">
            Define your repos once. Each workspace gets its own branches — no
            conflicts between agents, features, or experiments.
          </p>
        </div>
        <TerminalDemo />
      </section>

      <div className="h-px bg-[#172028]" />

      {/* MCP */}
      <section className="py-24 px-6 max-w-5xl mx-auto">
        <div className="grid md:grid-cols-2 gap-16 items-center">
          <div>
            <div className="text-[0.72rem] font-bold uppercase tracking-widest text-green-400 mb-3">
              AI Integration
            </div>
            <h2 className="text-3xl font-extrabold tracking-tight mb-4">
              Built-in MCP server for AI agents.
            </h2>
            <p className="text-[#7e8f9e] mb-6">
              hike includes a{" "}
              <a
                href="https://modelcontextprotocol.io"
                className="text-green-400 hover:text-green-300"
              >
                Model Context Protocol
              </a>{" "}
              server so AI assistants can manage workspaces autonomously.
            </p>
            <div className="flex flex-col gap-2.5 text-[#7e8f9e] text-sm">
              {[
                "Create & delete workspaces",
                "Pull, push, and check status across repos",
                "Sync repos to latest",
                "Works with Claude Code, Cursor, and any MCP client",
              ].map((text) => (
                <div key={text} className="flex items-center gap-2.5">
                  <svg
                    width="16"
                    height="16"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="#4ade80"
                    strokeWidth="2"
                    strokeLinecap="round"
                  >
                    <path d="M20 6L9 17l-5-5" />
                  </svg>
                  {text}
                </div>
              ))}
            </div>
          </div>
          <div className="rounded-xl border border-[#172028] bg-[#0a1016] overflow-hidden">
            <div className="flex items-center px-4 py-2.5 border-b border-[#172028] font-mono text-xs text-[#7e8f9e]">
              <span className="opacity-30 mr-3">● ● ●</span>
              claude_desktop_config.json
            </div>
            <pre className="p-5 font-mono text-sm leading-[1.8] overflow-x-auto">
              <span className="text-[#7e8f9e]">{"{"}</span>
              {"\n"}
              {"  "}
              <span className="text-sky-300">"mcpServers"</span>
              <span className="text-[#7e8f9e]">: {"{"}</span>
              {"\n"}
              {"    "}
              <span className="text-sky-300">"hike"</span>
              <span className="text-[#7e8f9e]">: {"{"}</span>
              {"\n"}
              {"      "}
              <span className="text-sky-300">"command"</span>
              <span className="text-[#7e8f9e]">: </span>
              <span className="text-green-400">"hike"</span>
              <span className="text-[#7e8f9e]">,</span>
              {"\n"}
              {"      "}
              <span className="text-sky-300">"args"</span>
              <span className="text-[#7e8f9e]">: [</span>
              <span className="text-green-400">"mcp"</span>
              <span className="text-[#7e8f9e]">],</span>
              {"\n"}
              {"      "}
              <span className="text-sky-300">"cwd"</span>
              <span className="text-[#7e8f9e]">: </span>
              <span className="text-green-400">"/path/to/projects"</span>
              {"\n"}
              {"    "}
              <span className="text-[#7e8f9e]">{"}"}</span>
              {"\n"}
              {"  "}
              <span className="text-[#7e8f9e]">{"}"}</span>
              {"\n"}
              <span className="text-[#7e8f9e]">{"}"}</span>
            </pre>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="text-center py-12 text-[#7e8f9e] text-sm border-t border-[#172028]">
        Not an official{" "}
        <a
          href="https://zuplo.com"
          className="text-[#7e8f9e] hover:text-green-400"
        >
          Zuplo
        </a>{" "}
        product. Free &amp; open source under the{" "}
        <a
          href="https://github.com/zuplo/hike/blob/main/LICENSE"
          className="text-[#7e8f9e] hover:text-green-400"
        >
          MIT License
        </a>
        .
      </footer>
    </div>
  );
};
