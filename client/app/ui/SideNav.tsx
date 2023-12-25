'use client';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faHome,
  faBriefcase,
  faBuilding,
} from '@fortawesome/free-solid-svg-icons';

import { usePathname } from 'next/navigation';
import Link from 'next/link';

import clsx from 'clsx';
import styles from '../styles/SideNav.module.css';

export default function SideNav() {
  const pathname = usePathname();

  return (
    <div id="sidenav" className={styles.sidenav}>
      <Link href="/" className={clsx({ [styles.active]: pathname === '/' })}>
        <FontAwesomeIcon icon={faHome} />
        <span>Home</span>
      </Link>
      <Link
        href="/applications"
        className={clsx({ [styles.active]: pathname === '/applications' })}
      >
        <FontAwesomeIcon icon={faBriefcase} />
        <span>Applications</span>
      </Link>
      <Link
        href="/companies"
        className={clsx({ [styles.active]: pathname === '/companies' })}
      >
        <FontAwesomeIcon icon={faBuilding} />
        <span>Companies</span>
      </Link>
    </div>
  );
}
