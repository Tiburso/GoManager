'use client';

import Searchbar from '@/app/ui/Searchbar';
import Button from '@/app/ui/Button';
import CompanyCards from '@/app/ui/companies/CompanyCards';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPlus } from '@fortawesome/free-solid-svg-icons';

import style from '@/app/styles/company/CompanyPage.module.css';

export default function CompanyPage() {
  // Change this to a lib call to get the companies from the API

  return (
    <div className="w-4/5 mx-auto">
      <div className={`${style.title}`}>
        {/* add class names for tailwind bigger text and centered and bold */}
        <h1 className="text-4xl text-left font-bold">Your Companies</h1>
        {/* onClick button should enable a modal to add a company */}
        <Button
          onClick={() => {
            console.log('clicked');
          }}
        >
          <FontAwesomeIcon icon={faPlus} className="fa fa-clone fa-xs mr-1" />
          Add company
        </Button>
      </div>
      <Searchbar type="text" placeholder="Search for company" />
      <CompanyCards />
    </div>
  );
}
