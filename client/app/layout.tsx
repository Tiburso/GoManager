import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';

import SideNav from './ui/SideNav';

// Check which fonts I should use for the application
const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  applicationName: 'GoManager',
  title: 'GoManager',
  description:
    'Next.js application used in combination with a Go API to manage job applications.',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <SideNav />
        {children}
      </body>
    </html>
  );
}
