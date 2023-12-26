import Searchbar from '@/app/ui/Searchbar';
import Button from '@/app/ui/Button';
import CompanyCard from '@/app/ui/companies/CompanyCard';

import { Company } from '@/app/lib/types';

import style from '@/app/styles//company/CompanyPage.module.css';

const CompanyData: Company[] = [
  {
    id: '1',
    name: 'Google',
    company_portal: 'https://www.google.com/',
  },
  {
    id: '2',
    name: 'Facebook',
    company_portal: 'https://www.facebook.com/',
  },
  {
    id: '3',
    name: 'Amazon',
    company_portal: 'https://www.amazon.com/',
  },
];

export default function CompanyPage() {
  // Change this to a lib call to get the companies from the API
  const companies = CompanyData;

  return (
    <div className="w-4/5 mx-auto">
      <div className={style.title}>
        {/* add class names for tailwind bigger text and centered and bold */}
        <h1 className="text-3xl text-left font-bold">Companies</h1>
        <Button>Add Company</Button>
      </div>
      <Searchbar type="text" placeholder="Search for company" />

      <div className="companiesList">
        {companies.map((company: Company) => (
          <CompanyCard key={company.id} company={company} />
        ))}
      </div>
    </div>
  );
}
