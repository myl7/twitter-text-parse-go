#ifndef TWITTER_TEXT_PARSE_GO
#define TWITTER_TEXT_PARSE_GO

#include <stdint.h>

struct TwTextParseRes {
    int32_t weighted_length;
    int32_t permillage;
    char is_valid;
    int32_t display_text_range_start;
    int32_t display_text_range_end;
    int32_t valid_text_range_start;
    int32_t valid_text_range_end;
};

const struct TwTextParseRes *tw_text_parse(const char *s);

#endif // TWITTER_TEXT_PARSE_GO
