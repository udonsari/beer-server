import openpyxl

# ************************************************************************

# TODO Main Code에 aroma4 추가하고, name_kor과 name_eng 구분 필요

ALPHABET_TO_COL = {
    'A': 'name_kor',
    'B': 'name_eng',
    'C': 'country',
    'D': 'abv',
    'E': 'style',
    'F': 'aroma1',
    'G': 'aroma2',
    'H': 'aroma3',
    'I': 'aroma4',
    'J': 'brewery'
}

ROW_MAX = 10
ROW_START = 2

def parse_beer(file_path):
    load_wb = openpyxl.load_workbook(file_path, data_only=True)
    load_ws = load_wb['Sheet1']

    beers = []

    for i in range(ROW_START, ROW_MAX):
        beer = {}
        for k, v in ALPHABET_TO_COL.items():
            beer[v] = load_ws[k + str(i)].value

        beer['abv'] *= 100
        beers.append(beer)

    for beer in beers:
        print(beer)

    return beers

    load_wb.close()

# ************************************************************************

TABLE_NAME = 'beer_info'

def build_insert_sql(beers):
    sql_str = """
    INSERT INTO {table_name}
    (
        name,
        brewery,
        abv,
        country,
        beer_style,
        aroma_list,
        image_url_list,
        thumbnail_image,
        rate_avg,
        review_count
    )
    VALUES""".format(
        table_name=TABLE_NAME
    )

    for idx, beer in enumerate(beers):
        if idx != 0:
            sql_str = sql_str + ','

        aroma_list_str = ''
        if beer['aroma1']:
            aroma_list_str += beer['aroma1']
        if beer['aroma2']:
            aroma_list_str += '___' + beer['aroma2']
        if beer['aroma3']:
            aroma_list_str += '___' + beer['aroma3']
        if beer['aroma4']:
            aroma_list_str += '___' + beer['aroma4']

        sql_str = sql_str + \
        """
        (
            \"{name}\",
            \"{brewery}\",
            {abv},
            \"{country}\",
            \"{style}\",
            \"{aroma_list}\",
            \"{image_url_list}\",
            \"{thumbnail_image}\",
            0,
            0
        )""".format(
            name=beer['name_eng'],
            brewery=beer['brewery'],
            abv=beer['abv'],
            country=beer['country'],
            style=beer['style'],
            aroma_list=aroma_list_str,
            image_url_list='TO_BE_FILLED', # TODO FILL
            thumbnail_image='TO_BE_FILLED', # TODO FILL
        )

    return sql_str

file_path = "./temp_beer_data.xlsx"
beers = parse_beer(file_path)
sql_str = build_insert_sql(beers)
print(sql_str)