import { hello } from "@scaffold/ui";

import { Button } from "@/components/ui/button";

console.log(hello());

function App() {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center">
      <div className="space-y-4 text-center">
        <h1 className="text-4xl font-bold text-foreground">Vite + Tailwind CSS + shadcn/ui</h1>

        <p className="text-muted-foreground">テスト用のページです</p>

        <div className="flex gap-4 justify-center">
          <Button>Default</Button>
          <Button variant="secondary">Secondary</Button>
          <Button variant="destructive">Destructive</Button>
          <Button variant="outline">Outline</Button>
          <Button variant="ghost">Ghost</Button>
        </div>

        <div className="mt-8 p-6 border rounded-lg">
          <h2 className="text-2xl font-semibold mb-2">カード例</h2>
          <p className="text-muted-foreground">
            Tailwind CSSのユーティリティクラスが正しく動作しています
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
