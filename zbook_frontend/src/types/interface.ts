interface MenuStruct {
  title: string;
  relative_path: string;
  isdir: boolean;
  sublayouts: Array<MenuStruct>;
  markdown_id: number;
}
interface Anchor {
  name: string;
  icon: string;
  url: string;
}
interface SubMenuProps {
  prefix: string;
  menus: Array<MenuStruct>;
  layer: number;
  pathname: string;
  locale: string;
  collapse: boolean;
}

interface SearchParams {
  [key: string]: string | undefined;
}

interface FooterSocial {
  name: string;
  icon: string;
  url: string;
}
export type { MenuStruct, SubMenuProps, SearchParams, Anchor, FooterSocial };
