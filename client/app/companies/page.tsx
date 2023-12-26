import Searchbar from '@/app/ui/Searchbar';
import CompanyCard from '@/app/ui/companies/CompanyCard';

import { Company } from '@/app/lib/types';

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
    <div>
      <h1>Your Companies</h1>
      <Searchbar />
      <div>
        <div className="companiesList">
          {companies.map((company: Company) => (
            <CompanyCard key={company.id} company={company} />
          ))}
        </div>
      </div>
    </div>
  );
}
