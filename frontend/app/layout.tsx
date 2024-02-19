import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { NumberProvider } from "./context/numbers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Numetric Flow",
  description: "Hello GoSolve! My name is Numertric Flow and I'm a technical interview app ;)",
  author: "Mateusz Solnica"
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <NumberProvider>  
          {children}
        </NumberProvider>
      </body>
    </html>
  );
}
