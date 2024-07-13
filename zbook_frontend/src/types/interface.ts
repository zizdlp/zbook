interface MenuStruct {
  title: string;
  relative_path: string;
  isdir: boolean;
  sublayouts: Array<MenuStruct>;
  markdown_id: number;
}
interface SubMenuProps {
  prefix: string;
  menus: Array<MenuStruct>;
  layer: number;
  pathname: string;
  locale: string;
}

interface SearchParams {
  [key: string]: string | undefined;
}

export type { MenuStruct, SubMenuProps, SearchParams };
