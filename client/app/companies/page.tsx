import { getCompanies } from '@/lib/companies';
import Companies from './companies';

// This component exists to do all the fetching of data from the API
// to do it via Server Side Rendering (SSR)
export default async function Page() {
  const companies = await getCompanies();

  return <Companies companies={companies} />;
}
